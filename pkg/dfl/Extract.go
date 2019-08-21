// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-try-get/pkg/gtg"
)

func parseExtractPath(path string) (int, int, int, int, error) {

	index_questionmark := -1
	index_period := -1
	index_start := -1
	index_end := -1

	for i, c := range path {
		if c == '?' {
			index_questionmark = i
			if i+1 < len(path) && path[i+1] == '.' {
				index_period = i + 1
				break
			} else {
				return 0, 0, 0, 0, errors.New("Invalid path " + path)
			}
		} else if c == '.' {
			index_period = i
			break
		} else if c == '[' {
			index_start = i
			leftsquarebrackets := 1
			rightsquarebrackets := 0
			for j, c := range path[i+1 : len(path)] {
				if c == '[' {
					leftsquarebrackets += 1
				} else if c == ']' {
					rightsquarebrackets += 1
					if leftsquarebrackets == rightsquarebrackets {
						index_end = j
						break
					}
				}
			}
			break
		}
	}

	return index_questionmark, index_period, index_start, index_end, nil
}

func checkStartIndex(obj interface{}, start int) error {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice || t.Kind() == reflect.String {
		s := reflect.ValueOf(obj)
		if start >= s.Len() {
			return errors.New("start index " + fmt.Sprint(start) + " is greater than or equal to the length of the object " + fmt.Sprint(s.Len()))
		}
	}
	return nil
}

func checkStartAndEndIndex(obj interface{}, start int, end int) error {
	t := reflect.TypeOf(obj)
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice || t.Kind() == reflect.String {
		s := reflect.ValueOf(obj)
		if end >= s.Len() {
			return errors.New("end index " + fmt.Sprint(end) + " greater than or equal to the length of the object " + fmt.Sprint(s.Len()))
		} else if start > end {
			return errors.New("start index " + fmt.Sprint(start) + " is greater than end index " + fmt.Sprint(end) + " for object")
		}
	}
	return nil
}

// Extract is a function to extract a value from an object.
// Extract supports a standard dot (.) and null-safe (?.) indexing.
// Extract also supports wildcard indexing using *.
// Extract also support array indexing, including [A], [A:B], [A:], and [:B].
func Extract(path string, obj interface{}, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	index_questionmark, index_period, slice_start_index, slice_end_index, err := parseExtractPath(path)
	if err != nil {
		return Null{}, errors.Wrap(err, "error parsing extract path")
	}

	if index_period != -1 {
		if index_questionmark != -1 {
			key := path[0:index_questionmark]
			remainder := path[index_period+1 : len(path)]

			t := reflect.TypeOf(obj)

			if t.Kind() == reflect.Map {
				m := reflect.ValueOf(obj)
				if m.Len() == 0 {
					return Null{}, nil
				}
				if key == "*" {
					values := reflect.MakeSlice(reflect.SliceOf(t.Elem()), 0, 0)
					for _, k := range m.MapKeys() {
						v, err := Extract(remainder, m.MapIndex(k).Interface(), vars, ctx, funcs, quotes)
						if err != nil {
							return v, err
						}
						values = reflect.Append(values, reflect.ValueOf(v))
					}
					return values.Interface(), nil
				}
				value := m.MapIndex(reflect.ValueOf(key))
				if !value.IsValid() {
					return Null{}, nil
				}
				return Extract(remainder, value.Interface(), vars, ctx, funcs, quotes)
			}

			switch o := obj.(type) {
			case *Context:
				if o.Has(key) {
					value := o.Get(key)
					if value == nil {
						return Null{}, nil
					}
					return Extract(remainder, value, vars, ctx, funcs, quotes)
				}
				return Null{}, nil
			}
			return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())

		} else {
			key := path[0:index_period]
			remainder := path[index_period+1 : len(path)]

			t := reflect.TypeOf(obj)

			if t.Kind() == reflect.Map {
				m := reflect.ValueOf(obj)
				if m.Len() == 0 {
					return Null{}, errors.New("value " + key + " is null.")
				}
				if key == "*" {
					values := reflect.MakeSlice(reflect.SliceOf(t.Elem()), 0, 0)
					for _, k := range m.MapKeys() {
						v, err := Extract(remainder, m.MapIndex(k).Interface(), vars, ctx, funcs, quotes)
						if err != nil {
							return v, err
						}
						values = reflect.Append(values, reflect.ValueOf(v))
					}
					return values.Interface(), nil
				}
				value := m.MapIndex(reflect.ValueOf(key))
				if !value.IsValid() {
					return Null{}, errors.New("value " + key + " is null.")
				}
				return Extract(remainder, value.Interface(), vars, ctx, funcs, quotes)
			}

			switch o := obj.(type) {
			case *Context:
				if key == "*" {
					if o.Len() == 0 {
						return Null{}, errors.New("value " + key + " is null.")
					}
					values := o.Values()
					results := make([]interface{}, 0, len(values))
					for _, v := range values {
						r, err := Extract(remainder, v, vars, ctx, funcs, quotes)
						if err != nil {
							return r, err
						}
						results = append(results, r)
					}
					return results, nil
				} else if o.Has(key) {
					value := o.Get(key)
					if value == nil {
						return Null{}, errors.New("value " + key + " is null.")
					}
					return Extract(remainder, value, vars, ctx, funcs, quotes)
				}
				return Null{}, errors.New("value " + key + " is null.")
			}
			return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())

		}
	} else if slice_start_index != -1 && slice_end_index != -1 {
		if slice_start_index == 0 {
			remainder := path[slice_end_index+2:]
			pair := strings.Split(path[1:slice_end_index+1], ":")
			if len(pair) == 2 {
				start := 0
				if len(pair[0]) > 0 {

					_, s, err := ParseCompileEvaluateInt(pair[0], vars, ctx, funcs, quotes)
					if err != nil {
						return Null{}, errors.New("slice start \"" + pair[0] + "\" is invalid")
					}

					err = checkStartIndex(obj, s)
					if err != nil {
						return Null{}, err
					}
					start = s
				}

				if len(pair[1]) > 0 {
					_, end, err := ParseCompileEvaluateInt(pair[1], vars, ctx, funcs, quotes)
					if err != nil {
						return Null{}, errors.New("slice end \"" + pair[1] + "\" is invalid")
					}

					err = checkStartAndEndIndex(obj, start, end)
					if err != nil {
						return Null{}, err
					}

					t := reflect.TypeOf(obj)
					if t.Kind() == reflect.Slice || t.Kind() == reflect.String {
						v := reflect.ValueOf(obj).Slice(start, end)
						if len(remainder) > 0 {
							return Extract(remainder, v.Interface(), vars, ctx, funcs, quotes)
						}
						return v.Interface(), nil
					}

					if o, ok := obj.(grw.ByteReadCloser); ok {
						return o.ReadRange(start, end-1)
					}

				} else {

					t := reflect.TypeOf(obj)
					if t.Kind() == reflect.Slice || t.Kind() == reflect.String {
						s := reflect.ValueOf(obj)
						if start >= s.Len() {
							return Null{}, errors.New("slice start index " + fmt.Sprint(start) + " greater than or equal to the length of the slice " + fmt.Sprint(s.Len()))
						}
						return s.Slice(start, s.Len()).Interface(), nil
					}

					if _, ok := obj.(grw.ByteReadCloser); ok {
						return make([]byte, 0), errors.New("Reader cannot evaluate [start:]")
					}

				}

				return Null{}, errors.New("object \"" + fmt.Sprint(obj) + "\" (" + reflect.TypeOf(obj).String() + ") is not a slice.")

			} else if len(pair) == 1 {

				_, i, err := ParseCompileEvaluate(pair[0], vars, ctx, funcs, quotes)
				if err != nil {
					return Null{}, errors.Wrap(err, "slice index \""+pair[0]+"\" is invalid ")
				}

				t := reflect.TypeOf(obj)
				if t.Kind() == reflect.Array || t.Kind() == reflect.Slice || t.Kind() == reflect.String {

					slice_index := 0
					switch i.(type) {
					case int:
						slice_index = i.(int)
					default:
						return Null{}, errors.New("slice index \"" + pair[0] + "\" is invalid type " + fmt.Sprint(reflect.TypeOf(i)))
					}

					s := reflect.ValueOf(obj)
					if slice_index >= s.Len() {
						return Null{}, errors.New("slice index " + fmt.Sprint(slice_index) + " greater than or equal to the length of the slice " + fmt.Sprint(s.Len()))
					}
					if len(remainder) > 0 {
						return Extract(remainder, s.Index(slice_index).Interface(), vars, ctx, funcs, quotes)
					}

					return s.Index(slice_index).Interface(), nil

				} else if t.Kind() == reflect.Map {
					m := reflect.ValueOf(obj)
					if len(remainder) > 0 {
						if t.Key().Kind() == reflect.String {
							return Extract(remainder, m.MapIndex(reflect.ValueOf(fmt.Sprint(i))).Interface(), vars, ctx, funcs, quotes)
						} else {
							return Extract(remainder, m.MapIndex(reflect.ValueOf(i)).Interface(), vars, ctx, funcs, quotes)
						}
					} else {
						if t.Key().Kind() == reflect.String {
							return m.MapIndex(reflect.ValueOf(fmt.Sprint(i))).Interface(), nil
						} else {
							return m.MapIndex(reflect.ValueOf(i)).Interface(), nil
						}
					}
				}

				if o, ok := obj.(grw.ByteReadCloser); ok {
					slice_index := 0
					switch i.(type) {
					case int:
						slice_index = i.(int)
					default:
						return Null{}, errors.New("slice index \"" + pair[0] + "\" is invalid type " + fmt.Sprint(reflect.TypeOf(i)))
					}
					values, err := o.ReadRange(slice_index, slice_index)
					if err != nil {
						return make([]byte, 0), err
					}
					return values[0], nil
				}

				return Null{}, errors.New("object \"" + fmt.Sprint(obj) + "\" (" + reflect.TypeOf(obj).String() + ") is not a slice.")
			}
			return Null{}, errors.New("slice range \"" + (path[1:slice_end_index]) + "\" is invalid ")
		} else {
			key := path[0:slice_start_index]
			remainder := path[slice_start_index:len(path)]

			t := reflect.TypeOf(obj)
			if t.Kind() == reflect.Map {
				m := reflect.ValueOf(obj)
				v := m.MapIndex(reflect.ValueOf(key))
				if !v.IsValid() {
					return Null{}, errors.New("value " + key + " is not present.")
				}
				if v.IsNil() {
					return Null{}, errors.New("value " + key + " is nil.")
				}
				return Extract(remainder, v.Interface(), vars, ctx, funcs, quotes)
			}

			switch o := obj.(type) {
			case *Context:
				if o.Has(key) {
					value := o.Get(key)
					if value == nil {
						return Null{}, errors.New("value " + key + " is null.")
					}
					return Extract(remainder, value, vars, ctx, funcs, quotes)
				}
				return Null{}, errors.New("value " + key + " is null.")
			}
			return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())
		}
	}

	if obj == nil {
		return Null{}, nil
	}

	t := reflect.TypeOf(obj)

	if t.Kind() == reflect.Map {
		m := reflect.ValueOf(obj)
		if m.Len() == 0 {
			return Null{}, nil
		}
		if path == "*" {
			values := reflect.MakeSlice(reflect.SliceOf(t.Elem()), 0, 0)
			for _, k := range m.MapKeys() {
				values = reflect.Append(values, m.MapIndex(k))
			}
			return values.Interface(), nil
		}

		value := gtg.TryGet(obj, path, nil)
		if value == nil {
			return Null{}, nil
		}
		return value, nil

	} else if t.Kind() == reflect.Struct {

		if path == "*" {
			return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())
		}

		value := gtg.TryGet(obj, path, nil)
		if value == nil {
			return Null{}, nil
		}
		return value, nil
	}

	switch o := obj.(type) {
	case *Context:
		if o.Has(path) {
			value := o.Get(path)
			if value == nil {
				return Null{}, nil
			}
			return value, nil
		}
		return Null{}, nil
	}

	return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())

}
