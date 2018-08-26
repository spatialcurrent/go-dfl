// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-reader/reader"
)

func intersects(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 2 {
		return Null{}, errors.New("Invalid number of arguments to intersects " + fmt.Sprint(len(args)))
	}

	switch a := args[0].(type) {
	case Null:
		switch b := args[1].(type) {
		case Null:
			return true, nil
		case StringSet:
			return b.Len() == 0, nil
		case map[string]struct{}:
			return len(b) == 0, nil
		}
	case StringSet:
		switch b := args[1].(type) {
		case Null:
			return a.Len() == 0, nil
		case StringSet:
			return a.Intersects(b), nil
		case map[string]struct{}:
			return a.Intersects(b), nil
		}
	case map[string]struct{}:
		switch b := args[1].(type) {
		case Null:
			return len(a) == 0, nil
		case map[string]struct{}:
			return StringSet(a).Intersects(b), nil
		case StringSet:
			return StringSet(a).Intersects(b), nil
		}
	}

	return Null{}, errors.New("Invalid arguments for intersects function " + reflect.TypeOf(args[0]).String() + " , " + reflect.TypeOf(args[1]).String())

}

func flattenArray(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 1 {
		return Null{}, errors.New("Invalid number of arguments to flattenArray.")
	}

	switch arr := args[0].(type) {
	case []string:
		return arr, nil
	case []int:
		return arr, nil
	case []interface{}:
		output := make([]interface{}, 0)
		for _, x := range arr {
			switch y := x.(type) {
			case []string:
				for _, z := range y {
					output = append(output, z)
				}
			case []interface{}:
				for _, z := range y {
					output = append(output, z)
				}
			case []int:
				for _, z := range y {
					output = append(output, z)
				}
			}
		}
		return output, nil
	}

	return Null{}, errors.New("Invalid arguments for flattenArray function " + reflect.TypeOf(args[0]).String())

}

func arrayToSet(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 1 {
		return Null{}, errors.New("Invalid number of arguments to arrayToSet (" + fmt.Sprint(len(args)) + ").")
	}

	switch arr := args[0].(type) {
	case Null:
		return Null{}, nil
	case []string:
		set := map[string]struct{}{}
		for _, v := range arr {
			set[v] = struct{}{}
		}
		return set, nil
	case []int:
		set := map[int]struct{}{}
		for _, v := range arr {
			set[v] = struct{}{}
		}
		return set, nil
	case []interface{}:
		set := map[interface{}]struct{}{}
		for _, v := range arr {
			set[v] = struct{}{}
		}
		return set, nil
	}

	return Null{}, errors.New("Invalid arguments for arrayToSet function " + reflect.TypeOf(args[0]).String())

}

func setToArray(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 1 {
		return Null{}, errors.New("Invalid number of arguments to setToArray.")
	}

	switch set := args[0].(type) {
	case map[string]struct{}:
		arr := make([]string, 0, len(set))
		for v := range set {
			arr = append(arr, v)
		}
		return arr, nil
	case map[int]struct{}:
		arr := make([]int, 0, len(set))
		for v := range set {
			arr = append(arr, v)
		}
		return arr, nil
	case map[interface{}]struct{}:
		arr := make([]interface{}, 0, len(set))
		for v := range set {
			arr = append(arr, v)
		}
		return TryConvertArray(arr), nil
	}

	return Null{}, errors.New("Invalid arguments for setToArray function " + reflect.TypeOf(args[0]).String())

}

func prefix(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to prefix.")
	}

	switch lv := args[0].(type) {
	case *reader.Cache:
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
	case []byte:
		switch prefix := args[1].(type) {
		case []byte:
			if len(prefix) > len(lv) {
				return false, nil
			}
			for i, c := range prefix {
				if lv[i] != c {
					return false, nil
				}
			}
			return true, nil
		case string:
			prefix_bytes := []byte(prefix)
			if len(prefix_bytes) > len(lv) {
				return false, nil
			}
			for i, c := range prefix_bytes {
				if lv[i] != c {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	case string:
		switch prefix := args[1].(type) {
		case string:
			return strings.HasPrefix(lv, prefix), nil
		}
		return Null{}, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	}

	return 0, errors.New("Invalid arguments for prefix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())

}

func suffix(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to suffix.")
	}

	switch lv := args[0].(type) {
	case *reader.Cache:
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
			//s := []rune(string(data))
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
	case []byte:
		switch suffix := args[1].(type) {
		case []byte:
			if len(suffix) > len(lv) {
				return false, nil
			}
			for i, _ := range suffix {
				if lv[len(lv)-i-1] != suffix[len(suffix)-i-1] {
					return false, nil
				}
			}
			return true, nil
		}
		return Null{}, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	case string:
		switch suffix := args[1].(type) {
		case string:
			return strings.HasSuffix(lv, suffix), nil
		}
		return Null{}, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
	}

	return 0, errors.New("Invalid arguments for suffix function " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())

}

func sortArray(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

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
			iv, err := key.Evaluate(arr[i], funcs, quotes)
			if err != nil {
				return false
			}
			jv, err := key.Evaluate(arr[j], funcs, quotes)
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
}

func limitArray(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to limitArray.")
	}

	switch limit := args[1].(type) {
	case int:
		switch arr := args[0].(type) {
		case []interface{}:
			if limit > len(arr) {
				return arr, nil
			}
			return arr[:limit], nil
		case []map[string]interface{}:
			if limit > len(arr) {
				return arr, nil
			}
			return arr[:limit], nil
		case []string:
			if limit > len(arr) {
				return arr, nil
			}
			return arr[:limit], nil
		case []int:
			if limit > len(arr) {
				return arr, nil
			}
			return arr[:limit], nil
		case []float64:
			if limit > len(arr) {
				return arr, nil
			}
			return arr[:limit], nil
		}
	}

	return 0, errors.New("Invalid arguments for limitArray " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
}

func filterArray(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 2 && len(args) != 3 {
		return 0, errors.New("Invalid number of arguments to filter.")
	}

	if len(args) == 2 {
		switch arr := args[0].(type) {
		case []map[string]interface{}:
			switch exp := args[1].(type) {
			case bool:
				if exp {
					return arr, nil
				}
				return make([]map[string]interface{}, 0), nil
			case string:
				n, err := Parse(exp)
				if err != nil {
					return 0, errors.Wrap(err, "error parsing expression for filter()")
				}
				n = n.Compile()
				output_slice := make([]map[string]interface{}, 0)
				for _, m := range arr {
					valid, err := EvaluateBool(n, m, funcs, quotes)
					if err != nil {
						return 0, errors.Wrap(err, "Error evaluating object "+fmt.Sprint(m))
					}
					if valid {
						output_slice = append(output_slice, m)
					}
				}
				return output_slice, nil
			}
		case []interface{}:
			switch exp := args[1].(type) {
			case bool:
				if exp {
					return arr, nil
				}
				return make([]interface{}, 0), nil
			case string:
				n, err := Parse(exp)
				if err != nil {
					return 0, errors.Wrap(err, "error parsing expression for filter()")
				}
				n = n.Compile()
				output_slice := make([]interface{}, 0)
				for _, m := range arr {
					valid, err := EvaluateBool(n, m, funcs, quotes)
					if err != nil {
						return 0, errors.Wrap(err, "Error evaluating object "+fmt.Sprint(m))
					}
					if valid {
						output_slice = append(output_slice, m)
					}
				}
				return output_slice, nil
			}
		}
	} else if len(args) == 3 {
		switch arr := args[0].(type) {
		case []map[string]interface{}:
			switch exp := args[1].(type) {
			case string:
				switch max_count := args[2].(type) {
				case int:
					n, err := Parse(exp)
					if err != nil {
						return 0, errors.Wrap(err, "error parsing expression for filter()")
					}
					n = n.Compile()
					output_slice := make([]map[string]interface{}, 0)
					for _, m := range arr {
						valid, err := EvaluateBool(n, m, funcs, quotes)
						if err != nil {
							return 0, errors.Wrap(err, "Error evaluating object "+fmt.Sprint(m))
						}
						if valid {
							output_slice = append(output_slice, m)
						}
						if max_count != -1 && len(output_slice) == max_count {
							break
						}
					}
					return output_slice, nil
				}
			}
		}
	}

	return 0, errors.New("Invalid arguments for filterArray " + reflect.TypeOf(args[0]).String() + ", " + reflect.TypeOf(args[1]).String())
}

func histArray(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 && len(args) != 2 && len(args) != 3 {
		return 0, errors.New("Invalid number of arguments to histArray.")
	}

	if len(args) == 3 {

		switch dfl_value := args[2].(type) {
		case string:
			node_values, err := Parse(dfl_value)
			if err != nil {
				return 0, errors.Wrap(err, "error parsing expression for histArray")
			}
			node_values = node_values.Compile()
			switch dfl_key := args[1].(type) {
			case string:
				node_key, err := Parse(dfl_key)
				if err != nil {
					return 0, errors.Wrap(err, "error parsing expression for histArray")
				}
				node_key = node_key.Compile()
				switch arr := args[0].(type) {
				case []map[string]interface{}:
					if len(arr) == 0 {
						return NewCounter(), nil
					}
					counters := map[string]map[string]int{}
					for _, x := range arr {
						x_key, err := node_key.Evaluate(x, funcs, quotes)
						if err != nil {
							return 0, errors.Wrap(err, "error extracting value from array element in histArray")
						}
						switch x_key_str := x_key.(type) {
						case string:
							if _, ok := counters[x_key_str]; !ok {
								counters[x_key_str] = map[string]int{}
							}
							x_values, err := node_values.Evaluate(x, funcs, quotes)
							if err != nil {
								return 0, errors.Wrap(err, "error extracting value from array element in histArray")
							}
							switch x_values.(type) {
							case []interface{}:
								for _, x_value := range x_values.([]interface{}) {
									Counter(counters[x_key_str]).Increment(fmt.Sprint(x_value))
								}
							case []string:
								for _, x_value := range x_values.([]string) {
									Counter(counters[x_key_str]).Increment(x_value)
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
						return NewCounter(), nil
					}
					h := map[string]Counter{}
					for _, x := range arr {
						x_key, err := node_key.Evaluate(x, funcs, quotes)
						if err != nil {
							return 0, errors.Wrap(err, "error extracting value from array element in histArray")
						}
						switch x_key_str := x_key.(type) {
						case string:
							if _, ok := h[x_key_str]; !ok {
								h[x_key_str] = NewCounter()
							}
							x_values, err := node_values.Evaluate(x, funcs, quotes)
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
			n, err := Parse(exp)
			if err != nil {
				return 0, errors.Wrap(err, "error parsing expression for histArray")
			}
			n = n.Compile()
			switch arr := args[0].(type) {
			case []map[string]interface{}:
				if len(arr) == 0 {
					return NewCounter(), nil
				}
				counter := NewCounter()
				for _, x := range arr {
					y, err := n.Evaluate(x, funcs, quotes)
					if err != nil {
						return 0, errors.Wrap(err, "error extracting value from array element in histArray")
					}
					counter.Increment(fmt.Sprint(y))
				}
				return counter, nil
			case []interface{}:
				if len(arr) == 0 {
					return NewCounter(), nil
				}
				counter := NewCounter()
				for _, x := range arr {
					y, err := n.Evaluate(x, funcs, quotes)
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
				return NewCounter(), nil
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
		case []string:
			counter := NewCounter()
			for _, value := range arr {
				counter.Increment(value)
			}
			return counter, nil
		}

	}

	return 0, errors.New("Invalid arguments for histArray " + reflect.TypeOf(args[0]).String())
}

func topCounter(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 2 && len(args) != 3 {
		return 0, errors.New("Invalid number of arguments to topCounter.")
	}

	if len(args) == 2 {
		switch n := args[1].(type) {
		case int:
			switch c := args[0].(type) {
			case Counter:
				return c.Top(n, 0), nil
			case map[string]int:
				return Counter(c).Top(n, 0), nil
			}
		}
	} else if len(args) == 3 {
		switch min := args[2].(type) {
		case int:
			switch n := args[1].(type) {
			case int:
				switch c := args[0].(type) {
				case Counter:
					return c.Top(n, min), nil
				case map[string]int:
					return Counter(c).Top(n, min), nil
				}
			}
		}
	}

	return 0, errors.New("Invalid arguments for topCounter " + reflect.TypeOf(args[0]).String() + " , " + reflect.TypeOf(args[1]).String())
}

func mapArray(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to map.")
	}

	switch exp := args[1].(type) {
	case string:
		n, err := Parse(exp)
		if err != nil {
			return 0, errors.Wrap(err, "error parsing expression for map")
		}
		n = n.Compile()
		switch arr := args[0].(type) {
		case []map[string]interface{}:
			values := make([]interface{}, 0, len(arr))
			for _, x := range arr {
				y, err := n.Evaluate(x, funcs, quotes) //Extract(key, x)
				if err != nil {
					return 0, errors.Wrap(err, "error extracting value from array element in mapArray")
				}
				values = append(values, y)
			}
			return values, nil
		case []map[string]string:
			values := make([]string, 0, len(arr))
			for _, x := range arr {
				y, err := n.Evaluate(x, funcs, quotes)
				if err != nil {
					return 0, errors.Wrap(err, "error extracting value from array element in mapArray")
				}
				values = append(values, fmt.Sprint(y))
			}
			return values, nil
		case []interface{}:
			values := make([]interface{}, 0, len(arr))
			for _, x := range arr {
				y, err := n.Evaluate(x, funcs, quotes)
				if err != nil {
					return 0, errors.Wrap(err, "error extracting value from array element in mapArray")
				}
				values = append(values, y)
			}
			return values, nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	return 0, errors.New("Invalid key for map function")

}

func splitString(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to split.")
	}
	switch s := args[0].(type) {
	case string:
		switch delim := args[1].(type) {
		case string:
			return strings.Split(s, delim), nil
		}
	case Null:
		return s, nil
	}
	return 0, errors.New("Invalid arguments for splitString " + reflect.TypeOf(args[0]).String() + " , " + reflect.TypeOf(args[1]).String())
}

func trimString(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to split.")
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimSpace(a), nil
	case []byte:
		return []byte(strings.TrimSpace(string(a))), nil
	case *reader.Cache:
		b, err := a.ReadAll()
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error reading all bytes from *reader.Cache")
		}
		return []byte(strings.TrimSpace(string(b))), nil
	}

	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func trimStringLeft(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to ltrim.")
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimLeftFunc(a, unicode.IsSpace), nil
	case []byte:
		return []byte(strings.TrimLeftFunc(string(a), unicode.IsSpace)), nil
	case *reader.Cache:
		i := 0
		for i = 0; ; i++ {
			b, err := a.ReadAt(i)
			if err != nil {
				if err == io.EOF {
					return make([]byte, 0), nil
				} else {
					return make([]byte, 0), errors.Wrap(err, "error reading byte at position "+fmt.Sprint(i)+" in trimStringLeft")
				}
			}
			if !unicode.IsSpace(bytes.Runes([]byte{b})[0]) {
				break
			}
		}
		return reader.NewCacheWithContent(a.Reader, a.Content, i), nil
	}

	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func trimStringRight(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to rtrim.")
	}

	switch a := args[0].(type) {
	case string:
		return strings.TrimRightFunc(a, unicode.IsSpace), nil
	case []byte:
		return []byte(strings.TrimRightFunc(string(a), unicode.IsSpace)), nil
	case *reader.Cache:
		b, err := a.ReadAll()
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error reading all bytes from *reader.Cache")
		}
		return []byte(strings.TrimRightFunc(string(b), unicode.IsSpace)), nil
	}
	return "", errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func getLength(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to len.")
	}

	switch a := args[0].(type) {
	case string:
		return len(a), nil
	case []int:
		return len(a), nil
	case []string:
		return len(a), nil
	case []uint8:
		return len(a), nil
	case []float64:
		return len(a), nil
	case []interface{}:
		return len(a), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToBytes(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to len.")
	}

	switch a := args[0].(type) {
	case string:
		return []byte(a), nil
	case byte:
		return []byte{a}, nil
	case []byte:
		return a, nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToInt16(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to int16.")
	}
	switch a := args[0].(type) {
	case int:
		return int16(a), nil
	case int8:
		return int16(a), nil
	case int16:
		return a, nil
	case int32:
		return int16(a), nil
	case int64:
		return int16(a), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToInt32(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to int32.")
	}
	switch a := args[0].(type) {
	case int:
		return int32(a), nil
	case int16:
		return int32(a), nil
	case int32:
		return a, nil
	case int64:
		return int32(a), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToBigEndian(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to big-endian.")
	}
	switch a := args[0].(type) {
	case int:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, int64(a))
		return buf.Bytes(), err
	case int16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, a)
		return buf.Bytes(), err
	case int32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, a)
		return buf.Bytes(), err
	case int64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.BigEndian, a)
		return buf.Bytes(), err
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToLittleEndian(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to little-endian.")
	}
	switch a := args[0].(type) {
	case int:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, int64(a))
		return buf.Bytes(), err
	case int16:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, a)
		return buf.Bytes(), err
	case int32:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, a)
		return buf.Bytes(), err
	case int64:
		buf := new(bytes.Buffer)
		err := binary.Write(buf, binary.LittleEndian, a)
		return buf.Bytes(), err
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func repeat(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 2 {
		return 0, errors.New("Invalid number of arguments to repeat.")
	}
	switch value := args[0].(type) {
	case string:
		switch count := args[1].(type) {
		case int:
			return strings.Repeat(value, count), nil
		}
	case []byte:
		switch count := args[1].(type) {
		case int:
			return bytes.Repeat(value, count), nil
		}
	case byte:
		switch count := args[1].(type) {
		case int:
			return bytes.Repeat([]byte{value}, count), nil
		}
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}

func convertToString(funcs FunctionMap, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
	if len(args) != 1 {
		return 0, errors.New("Invalid number of arguments to len.")
	}
	switch a := args[0].(type) {
	case string:
		return a, nil
	case []byte:
		return string(a), nil
	case byte:
		return string([]byte{a}), nil
	case *reader.Cache:
		value, err := a.ReadAll()
		if err != nil {
			return "", errors.Wrap(err, "error reading all content from *reader.Cache in covertToString")
		}
		return string(value), nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
}
