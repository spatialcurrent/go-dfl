// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unicode"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
	"github.com/spatialcurrent/go-counter/pkg/counter"
	"github.com/spatialcurrent/go-reader-writer/pkg/cache"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

func toDict(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) > 3 {
		return Null{}, errors.New("Invalid number of arguments to toDict.")
	}

	t := reflect.TypeOf(args[0])
	if t.Kind() != reflect.Array && t.Kind() != reflect.Slice {
		return Null{}, errors.New("Invalid arguments for toDict function " + reflect.TypeOf(args[0]).String())
	}
	v := reflect.ValueOf(args[0])
	l := v.Len()
	if l == 0 {
		return map[string]interface{}{}, nil
	}

	validKeys := map[string]struct{}{}
	if len(args) == 2 {
		keys, err := af.ToStringSet.ValidateRun([]interface{}{args[1]})
		if err != nil {
			return Null{}, errors.New("Invalid arguments for toDict function " + reflect.TypeOf(args[1]).String())
		}
		validKeys = keys.(map[string]struct{})
	}

	m := map[interface{}]interface{}{}
	for i := 0; i < l; i++ {
		x := v.Index(i).Interface()
		t2 := reflect.TypeOf(x)
		if t2.Kind() != reflect.Array && t2.Kind() != reflect.Slice {
			return Null{}, errors.New("Invalid item for toDict function " + t2.String())
		}
		v2 := reflect.ValueOf(x)
		if v2.Len() != 2 {
			return Null{}, errors.New("Invalid length for item for toDict function " + fmt.Sprint(v2.Len()))
		}
		k := v2.Index(0).Interface()
		valid := true
		if len(validKeys) > 0 {
			if _, ok := validKeys[fmt.Sprint(k)]; !ok {
				valid = false
			}
		}
		if valid {
			m[k] = v2.Index(1).Interface()
		}
	}
	return m, nil
}

func prefix(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to prefix.")
	}

	if lv, ok := args[0].(io.ByteReadCloser); ok {
		switch prefix := args[1].(type) {
		case []byte:
			data, err := lv.ReadRange(0, len(prefix)-1)
			if err != nil {
				return false, nil
			}
			for i, c := range prefix {
				if data[i] != c {
					return false, nil
				}
			}
			return true, nil
		case string:
			data, err := lv.ReadRange(0, len(prefix)-1)
			if err != nil {
				return false, nil
			}
			s := []rune(string(data))
			if len(s) < len(prefix) {
				return false, nil
			}
			for i, c := range prefix {
				if s[i] != c {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	}

	return af.Prefix.ValidateRun(args)

}

func suffix(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to suffix.")
	}

	if lv, ok := args[0].(io.ByteReadCloser); ok {
		switch suffix := args[1].(type) {
		case []byte:
			data, err := lv.ReadAll()
			if err != nil {
				return false, nil
			}
			if len(suffix) > len(data) {
				return false, nil
			}
			for i, _ := range suffix {
				if data[len(data)-i-1] != suffix[len(suffix)-i-1] {
					return false, nil
				}
			}
			return true, nil
		case string:
			data, err := lv.ReadAll()
			if err != nil {
				return false, nil
			}
			s := string(data)
			if len(suffix) > len(s) {
				return false, nil
			}
			for i, _ := range suffix {
				if s[len(s)-i-1] != suffix[len(suffix)-i-1] {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	}

	return af.Suffix.ValidateRun(args)
}

/*
// no current uses for custom sorting, so just going to use simpler one from go-adaptive-functions
func sortArray(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 1 && len(args) != 2 && len(args) != 3 {
		return 0, errors.New("Invalid number of arguments to sortArray.")
	}

	s := reflect.ValueOf(args[0])
	if s.Kind() != reflect.Slice {
		return "", errors.New("Argument for sortArray is not kind slice.")
	}

	descending := false
	if len(args) == 3 {
		switch args[2].(type) {
		case bool:
			descending = args[2].(bool)
		default:
			return "", errors.New("argument ascending for sortArray is not a boolean but " + fmt.Sprint(reflect.TypeOf(args[1])))
		}
	}

	var arr interface{}
	switch args[0].(type) {
	case []interface{}:
		arr = TryConvertArray(args[0].([]interface{}))
	default:
		arr = args[0]
	}

	var key Node
	if len(args) >= 2 {
		switch exp := args[1].(type) {
		case string:
			if len(exp) > 0 {
				n, err := Parse(exp)
				if err != nil {
					return 0, errors.Wrap(err, "error parsing expression for key")
				}
				key = n.Compile()
			}
		default:
			return 0, errors.New("Invalid sort key (" + fmt.Sprint(reflect.TypeOf(args[2])) + ") for []interface{}")
		}
	}

	switch arr := arr.(type) {
	case []interface{}:
		if len(arr) == 0 {
			return arr, nil
		}
		if key == nil {
			return 0, errors.New("Cannot sort []interface{} without a key, because no natural sort order.")
		}
		sort.SliceStable(arr, func(i int, j int) bool {
			_, iv, err := key.Evaluate(vars, arr[i], funcs, quotes)
			if err != nil {
				return false
			}
			_, jv, err := key.Evaluate(vars, arr[j], funcs, quotes)
			if err != nil {
				return false
			}
			switch iv.(type) {
			case string:
				switch jv.(type) {
				case string:
					if descending {
						return strings.Compare(iv.(string), jv.(string)) > 0
					}
					return strings.Compare(iv.(string), jv.(string)) < 0
				}
			}
			r, err := CompareNumbers(iv, jv)
			if err != nil {
				return false
			}
			if descending {
				return r > 0
			}
			return r < 0
		})
		return arr, nil
	case []string:
		if descending {
			sort.Sort(sort.Reverse(sort.StringSlice(arr)))
		} else {
			sort.Strings(arr)
		}
		return arr, nil
	case []int:
		if descending {
			sort.Sort(sort.Reverse(sort.IntSlice(arr)))
		} else {
			sort.Ints(arr)
		}
		return arr, nil
	case []float64:
		if descending {
			sort.Sort(sort.Reverse(sort.Float64Slice(arr)))
		} else {
			sort.Float64s(arr)
		}
		return arr, nil
	}

	return 0, errors.New("Invalid arguments for sortArray " + reflect.TypeOf(args[0]).String())
}*/

func filterArray(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 2 && len(args) != 3 {
		return 0, errors.New("Invalid number of arguments to filter.")
	}

	arrayType := reflect.TypeOf(args[0])

	if arrayType.Kind() != reflect.Array && arrayType.Kind() != reflect.Slice {
		return 0, errors.New("Invalid arguments for filterArray " + arrayType.String() + ", " + reflect.TypeOf(args[1]).String())
	}

	if t := reflect.TypeOf(args[1]); t.Kind() == reflect.Bool {
		if args[1].(bool) {
			return args[0], nil
		}
		return reflect.MakeSlice(t, 0, 0).Interface(), nil
	}

	outputLimit := -1
	if len(args) == 3 {
		if t := reflect.TypeOf(args[2]); t.Kind() != reflect.Int {
			return 0, errors.New("Invalid max count for filterArray " + t.String())
		}
		outputLimit = args[2].(int)
	}

	var node Node
	t := reflect.TypeOf(args[1])
	if t.Kind() == reflect.String {
		n, err := ParseCompile(args[1].(string))
		if err != nil {
			return 0, errors.Wrap(err, "error parsing expression for filter()")
		}
		node = n
	} else if n, ok := args[1].(Node); ok {
		node = n
	} else {
		return 0, errors.New("Invalid arguments for filterArray " + reflect.TypeOf(args[0]).String() + ", " + fmt.Sprint(t))
	}

	originalSlice := reflect.ValueOf(args[0])
	originalLength := originalSlice.Len()

	output_slice := reflect.MakeSlice(arrayType, 0, 0)
	for i := 0; i < originalLength; i++ {
		m := originalSlice.Index(i).Interface()
		_, valid, err := node.Evaluate(vars, m, funcs, quotes)
		if err != nil {
			return 0, errors.Wrap(err, "error evaluating object "+fmt.Sprint(m))
		}
		if reflect.TypeOf(valid).Kind() == reflect.Bool && valid.(bool) {
			output_slice = reflect.Append(output_slice, reflect.ValueOf(m))
		}
		if outputLimit >= 0 && output_slice.Len() == outputLimit {
			break
		}
	}

	return output_slice.Interface(), nil

}

func groupArray(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) < 2 || len(args) > 6 {
		return 0, errors.New("Invalid number of arguments to group.")
	}

	arrayType := reflect.TypeOf(args[0])

	if arrayType.Kind() != reflect.Array && arrayType.Kind() != reflect.Slice {
		return 0, errors.New("Invalid arguments for groupArray " + arrayType.String() + ", " + reflect.TypeOf(args[1]).String())
	}

	if t := reflect.TypeOf(args[1]); t.Kind() == reflect.Bool {
		if args[1].(bool) {
			return args[0], nil
		}
		return reflect.MakeSlice(t, 0, 0).Interface(), nil
	}

	outputLimit := -1
	if len(args) == 4 {
		if i, ok := args[3].(int); ok {
			outputLimit = i
		} else {
			return 0, errors.New("Invalid max count for groupArray " + reflect.TypeOf(args[3]).String())
		}
	}

	var nodeAccumulator Node
	var initialValue interface{}

	if len(args) >= 5 {
		t := reflect.TypeOf(args[1])
		if t.Kind() == reflect.String {
			n, err := ParseCompile(args[4].(string))
			if err != nil {
				return 0, errors.Wrap(err, "error parsing aggregation for groupArray")
			}
			nodeAccumulator = n
		} else if n, ok := args[4].(Node); ok {
			nodeAccumulator = n
		} else {
			return 0, errors.New("Invalid arguments for filterArray " + reflect.TypeOf(args[0]).String() + ", " + fmt.Sprint(t))
		}

		if len(args) >= 6 {
			_, iv, err := ParseCompileEvaluate(args[5].(string), vars, map[string]interface{}{}, funcs, quotes)
			if err != nil {
				return 0, errors.Wrap(err, "error parsing initial value for groupArray")
			}
			initialValue = iv
		}
	}

	var nodeKeys Node
	t := reflect.TypeOf(args[1])
	if t.Kind() == reflect.String {
		n, err := ParseCompile(args[1].(string))
		if err != nil {
			return 0, errors.Wrap(err, "error parsing expression for filter()")
		}
		nodeKeys = n
	} else if n, ok := args[1].(Node); ok {
		nodeKeys = n
	} else {
		return 0, errors.New("Invalid arguments for filterArray " + reflect.TypeOf(args[0]).String() + ", " + fmt.Sprint(t))
	}

	var nodeOutput Node
	if len(args) >= 3 {
		if str, ok := args[2].(string); ok {
			n, err := ParseCompile(str)
			if err != nil {
				return 0, errors.Wrap(err, "error parsing expression for filter()")
			}
			nodeOutput = n
		} else if n, ok := args[2].(Node); ok {
			nodeOutput = n
		} else {
			return 0, errors.New("Invalid arguments for filterArray " + reflect.TypeOf(args[2]).String() + ", " + fmt.Sprint(t))
		}
	}

	originalSlice := reflect.ValueOf(args[0])
	originalLength := originalSlice.Len()

	if _, ok := nodeKeys.(Lengther); !ok {
		return 0, errors.New("cannot get length from node")
	}

	outputMap := CreateGroups(nodeKeys.(Lengther).Len())
	for i := 0; i < originalLength; i++ {
		obj := originalSlice.Index(i)
		_, keys, err := EvaluateArray(nodeKeys, vars, obj.Interface(), funcs, quotes)
		if err != nil {
			return 0, errors.Wrap(err, "error evaluating object "+fmt.Sprint(obj.Interface()))
		}
		currentMap := reflect.ValueOf(outputMap)
		keysValue := reflect.ValueOf(keys)
		for j := 0; j < keysValue.Len(); j++ {
			key := fmt.Sprint(keysValue.Index(j).Interface())
			next := currentMap.MapIndex(reflect.ValueOf(key))
			if j < keysValue.Len()-1 {
				if !next.IsValid() {
					next = reflect.MakeMap(currentMap.Type().Elem())
					currentMap.SetMapIndex(reflect.ValueOf(key), next)
				}
				currentMap = next
			} else {

				var outputObject reflect.Value
				if nodeOutput != nil {
					_, o, err := nodeOutput.Evaluate(vars, obj.Interface(), funcs, quotes)
					if err != nil {
						return 0, errors.Wrap(err, "error evaluating output object "+fmt.Sprint(obj.Interface()))
					}
					outputObject = reflect.ValueOf(o)
				} else {
					outputObject = obj
				}

				if next.IsValid() {
					if nodeAccumulator != nil {
						_, accumulatedValue, err := nodeAccumulator.Evaluate(vars, []interface{}{next.Interface(), outputObject.Interface()}, funcs, quotes)
						if err != nil {
							return 0, errors.Wrap(err, "error evaluating accumuator")
						}
						currentMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(accumulatedValue))
					} else {
						currentMap.SetMapIndex(reflect.ValueOf(key), reflect.Append(reflect.ValueOf(next.Interface()), outputObject))
					}
				} else {
					if initialValue != nil && nodeAccumulator != nil {
						_, accumulatedValue, err := nodeAccumulator.Evaluate(vars, []interface{}{initialValue, outputObject.Interface()}, funcs, quotes)
						if err != nil {
							return 0, errors.Wrap(err, "error evaluating accumuator")
						}
						currentMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(accumulatedValue))
					} else {
						currentMap.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf([]interface{}{outputObject.Interface()}))
					}
				}
			}
		}

		if outputLimit >= 0 && (i+1) == outputLimit {
			break
		}
	}

	return outputMap, nil

}

func histArray(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 && len(args) != 2 && len(args) != 3 {
		return 0, errors.New("Invalid number of arguments to histArray.")
	}

	if len(args) == 3 {

		switch dfl_value := args[2].(type) {
		case string:
			node_values, err := ParseCompile(dfl_value)
			if err != nil {
				return 0, errors.Wrap(err, "error parsing expression for histArray")
			}
			switch dfl_key := args[1].(type) {
			case string:
				node_key, err := ParseCompile(dfl_key)
				if err != nil {
					return 0, errors.Wrap(err, "error parsing expression for histArray")
				}
				switch arr := args[0].(type) {
				case []map[string]interface{}:
					if len(arr) == 0 {
						return counter.New(), nil
					}
					counters := map[string]map[string]int{}
					for _, x := range arr {
						_, x_key, err := node_key.Evaluate(vars, x, funcs, quotes)
						if err != nil {
							return 0, errors.Wrap(err, "error extracting value from array element in histArray")
						}
						switch x_key_str := x_key.(type) {
						case string:
							if _, ok := counters[x_key_str]; !ok {
								counters[x_key_str] = map[string]int{}
							}
							_, x_values, err := node_values.Evaluate(vars, x, funcs, quotes)
							if err != nil {
								return 0, errors.Wrap(err, "error extracting value from array element in histArray")
							}
							switch x_values.(type) {
							case StringSet:
								for x_value, _ := range x_values.(StringSet) {
									counter.Counter(counters[x_key_str]).Increment(x_value)
								}
							case map[string]struct{}:
								for x_value, _ := range x_values.(map[string]struct{}) {
									counter.Counter(counters[x_key_str]).Increment(x_value)
								}
							case []interface{}:
								for _, x_value := range x_values.([]interface{}) {
									counter.Counter(counters[x_key_str]).Increment(fmt.Sprint(x_value))
								}
							case []string:
								for _, x_value := range x_values.([]string) {
									counter.Counter(counters[x_key_str]).Increment(x_value)
								}
							default:
								return 0, errors.Wrap(err, "invalid histogram values "+fmt.Sprint(reflect.TypeOf(x_values)))
							}
						default:
							return 0, errors.Wrap(err, "invalid histogram key "+fmt.Sprint(reflect.TypeOf(x_key)))
						}
					}
					return counters, nil
				case []interface{}:
					if len(arr) == 0 {
						return counter.New(), nil
					}
					h := map[string]counter.Counter{}
					for _, x := range arr {
						_, x_key, err := node_key.Evaluate(vars, x, funcs, quotes)
						if err != nil {
							return 0, errors.Wrap(err, "error extracting value from array element in histArray")
						}
						switch x_key_str := x_key.(type) {
						case string:
							if _, ok := h[x_key_str]; !ok {
								h[x_key_str] = counter.New()
							}
							_, x_values, err := node_values.Evaluate(vars, x, funcs, quotes)
							if err != nil {
								return 0, errors.Wrap(err, "error extracting value from array element in histArray")
							}
							switch x_values.(type) {
							case []interface{}:
								for _, x_value := range x_values.([]interface{}) {
									h[x_key_str].Increment(fmt.Sprint(x_value))
								}
							case []string:
								for _, x_value := range x_values.([]string) {
									h[x_key_str].Increment(x_value)
								}
							default:
								return 0, errors.Wrap(err, "invalid histogram values "+fmt.Sprint(reflect.TypeOf(x_values)))
							}
						default:
							return 0, errors.Wrap(err, "invalid histogram key "+fmt.Sprint(reflect.TypeOf(x_key)))
						}
					}
					return h, nil
				}
			}
		}

	} else if len(args) == 2 {

		switch exp := args[1].(type) {
		case string:
			n, err := ParseCompile(exp)
			if err != nil {
				return 0, errors.Wrap(err, "error parsing expression for histArray")
			}
			switch arr := args[0].(type) {
			case []map[string]interface{}:
				if len(arr) == 0 {
					return counter.New(), nil
				}
				counter := counter.New()
				for _, x := range arr {
					_, y, err := n.Evaluate(vars, x, funcs, quotes)
					if err != nil {
						return 0, errors.Wrap(err, "error extracting value from array element in histArray")
					}
					counter.Increment(fmt.Sprint(y))
				}
				return counter, nil
			case []interface{}:
				if len(arr) == 0 {
					return counter.New(), nil
				}
				counter := counter.New()
				for _, x := range arr {
					_, y, err := n.Evaluate(vars, x, funcs, quotes)
					if err != nil {
						return 0, errors.Wrap(err, "error extracting value from array element in histArray")
					}
					counter.Increment(fmt.Sprint(y))
				}
				return counter, nil
			default:
				return 0, errors.Wrap(err, "invalid histogram values "+fmt.Sprint(reflect.TypeOf(arr)))
			}
		}

	} else if len(args) == 1 {
		var values interface{}

		switch arr := args[0].(type) {
		case []interface{}:
			if len(arr) == 0 {
				return counter.New(), nil
			}
			switch arr2 := TryConvertArray(arr).(type) {
			case []interface{}:
				values := make([]string, 0, len(arr2))
				for _, x := range arr2 {
					values = append(values, fmt.Sprint(x))
				}
			default:
				values = arr2
			}
		default:
			values = arr
		}

		switch arr := values.(type) {
		case map[string]interface{}:
			counter := counter.New()
			for key, _ := range arr {
				counter.Increment(key)
			}
			return counter, nil
		case StringSet:
			counter := counter.New()
			for key, _ := range arr {
				counter.Increment(key)
			}
			return counter, nil
		case []string:
			counter := counter.New()
			for _, value := range arr {
				counter.Increment(value)
			}
			return counter, nil
		}

	}

	return 0, errors.New("Invalid arguments for histArray " + reflect.TypeOf(args[0]).String())
}

func mapArray(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to map.")
	}

	var node Node
	t := reflect.TypeOf(args[1])
	if t.Kind() == reflect.String {
		n, err := ParseCompile(strings.TrimSpace(args[1].(string)))
		if err != nil {
			return 0, errors.Wrap(err, "error parsing expression for map("+(args[1].(string))+")")
		}
		node = n
	} else {
		return 0, errors.New("Invalid arguments for mapArray " + reflect.TypeOf(args[0]).String() + ", " + fmt.Sprint(t))
	}

	t = reflect.TypeOf(args[0])
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		original := reflect.ValueOf(args[0])
		if original.Len() == 0 {
			// Always convert to []interface{}, since some downstream functions might assume this
			return make([]interface{}, 0), nil
		}
		length := original.Len()
		values := make([]interface{}, 0, length)
		for i := 0; i < length; i++ {
			_, y, err := node.Evaluate(vars, original.Index(i).Interface(), funcs, quotes)
			if err != nil {
				return 0, errors.Wrap(err, fmt.Sprint("error evaluating value for element:", original.Index(i).Interface()))
			}
			values = append(values, y)
		}
		return values, nil
	} else if t.Kind() == reflect.Map {
		original := reflect.ValueOf(args[0])
		if original.Len() == 0 {
			return args[0], nil
		}
		it := reflect.TypeOf((*interface{})(nil)).Elem()
		values := reflect.MakeMap(reflect.MapOf(t.Key(), it))
		for _, k := range original.MapKeys() {
			_, y, err := node.Evaluate(vars, original.MapIndex(k).Interface(), funcs, quotes)
			if err != nil {
				return 0, errors.Wrap(err, fmt.Sprint("error extracting value from map element in map:", original.MapIndex(k).Interface()))
			}
			values.SetMapIndex(k, reflect.ValueOf(y))
		}
		return values.Interface(), nil
	}
	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())

}

func trimString(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to split.")
	}

	if a, ok := args[0].(io.ByteReadCloser); ok {
		b, err := a.ReadAll()
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error reading all bytes from *reader.Cache")
		}
		return []byte(strings.TrimSpace(string(b))), nil
	}

	switch a := args[0].(type) {
	case []byte:
		return []byte(strings.TrimSpace(string(a))), nil
	}

	if err := af.Trim.Validate(args); err != nil {
		return Null{}, errors.Wrap(err, "Invalid arguments")
	}
	return af.Trim.Run(args)
}

func trimStringLeft(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to ltrim.")
	}

	if a, ok := args[0].(io.ByteReadCloser); ok {
		content := make([]byte, 0)
		i := 0
		for i = 0; ; i++ {
			b := make([]byte, 1)
			_, err := a.ReadAt(b, int64(i))
			if err != nil {
				if err == io.EOF {
					return make([]byte, 0), nil
				} else {
					return make([]byte, 0), errors.Wrap(err, "error reading byte at position "+fmt.Sprint(i)+" in trimStringLeft")
				}
			}
			if !unicode.IsSpace(bytes.Runes(b)[0]) {
				content = append(content, b...)
				break
			}
		}
		return cache.NewCacheWithContent(a, &content, i), nil
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimLeftFunc(a, unicode.IsSpace), nil
	case []byte:
		return []byte(strings.TrimLeftFunc(string(a), unicode.IsSpace)), nil
	}

	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func trimStringRight(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to rtrim.")
	}

	if a, ok := args[0].(io.ByteReadCloser); ok {
		b, err := a.ReadAll()
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error reading all bytes from *reader.Cache")
		}
		return []byte(strings.TrimRightFunc(string(b), unicode.IsSpace)), nil
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimRightFunc(a, unicode.IsSpace), nil
	case []byte:
		return []byte(strings.TrimRightFunc(string(a), unicode.IsSpace)), nil
	}

	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToString(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to convertToString.")
	}

	if a, ok := args[0].(io.ByteReadCloser); ok {
		value, err := a.ReadAll()
		if err != nil {
			return "", errors.Wrap(err, "error reading all content from *reader.Cache in covertToString")
		}
		return string(value), nil
	}

	if err := af.ToString.Validate(args); err != nil {
		return Null{}, errors.Wrap(err, "Invalid arguments")
	}
	return af.ToString.Run(args)
}
