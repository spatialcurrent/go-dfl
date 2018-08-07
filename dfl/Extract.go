// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// Extract is a function to extract a value from an object.
// Extract supports a standard dot (.) and null-safe (?.) indexing.
// Extract also support array indexing, including [A], [A:B], [A:], and [:B].
func Extract(path string, obj interface{}) (interface{}, error) {

	index_questionmark := -1
	index_period := -1
	slice_start_index := -1
	slice_end_index := -1

	for i, c := range path {
		if c == '?' {
			index_questionmark = i
			if i+1 < len(path) && path[i+1] == '.' {
				index_period = i + 1
				break
			} else {
				return Null{}, errors.New("Invalid path " + path)
			}
		} else if c == '.' {
			index_period = i
			break
		} else if c == '[' {
			slice_start_index = i
			for j, c := range path[i+1 : len(path)] {
				if c == ']' {
					slice_end_index = j
					break
				}
			}
			break
		}
	}

	if index_period != -1 {
		if index_questionmark != -1 {
			key := path[0:index_questionmark]
			remainder := path[index_period+1 : len(path)]

			switch obj.(type) {
			case Context:
				if value, ok := obj.(Context)[key]; ok {
					if value == nil {
						return Null{}, nil
					}
					return Extract(remainder, value)
				}
				return Null{}, nil
			case map[string]interface{}:
				if value, ok := obj.(map[string]interface{})[key]; ok {
					if value == nil {
						return Null{}, nil
					}
					return Extract(remainder, value)
				}
				return Null{}, nil
			}
			return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())

		} else {
			key := path[0:index_period]
			remainder := path[index_period+1 : len(path)]

			switch obj.(type) {
			case Context:
				if value, ok := obj.(Context)[key]; ok {
					if value == nil {
						return Null{}, errors.New("value " + key + " is null.")
					}
					return Extract(remainder, value)
				}
				return Null{}, errors.New("value " + key + " is null.")
			case map[string]interface{}:
				if value, ok := obj.(map[string]interface{})[key]; ok {
					if value == nil {
						return Null{}, errors.New("value " + key + " is null.")
					}
					return Extract(remainder, value)
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
					i, err := strconv.Atoi(pair[0])
					if err != nil {
						return Null{}, errors.New("slice start \"" + pair[0] + "\" is invalid ")
					}
					start = i
				}

				if len(pair[1]) > 0 {
					end, err := strconv.Atoi(pair[1])
					if err != nil {
						return Null{}, errors.New("slice end \"" + pair[1] + "\" is invalid ")
					}

					switch o := obj.(type) {
					case string:
						return o[start:end], nil
					case []byte:
						return o[start:end], nil
					case []int:
						return o[start:end], nil
					case []string:
						return o[start:end], nil
					case []map[string]interface{}:
						if len(remainder) > 0 {
							return Extract(remainder, o[start:end])
						}
						return o[start:end], nil
					case []map[string]string:
						if len(remainder) > 0 {
							return Extract(remainder, o[start:end])
						}
						return o[start:end], nil
					}

				} else {

					switch o := obj.(type) {
					case string:
						return o[start:], nil
					case []byte:
						return o[start:], nil
					case []int:
						return o[start:], nil
					case []string:
						return o[start:], nil
					case []map[string]interface{}:
						if len(remainder) > 0 {
							return Extract(remainder, o[start:])
						}
						return o[start:], nil
					case []map[string]string:
						if len(remainder) > 0 {
							return Extract(remainder, o[start:])
						}
						return o[start:], nil
					}

				}

				return Null{}, errors.New("object \"" + fmt.Sprint(obj) + "\" (" + reflect.TypeOf(obj).String() + ") is not a slice.")

			} else if len(pair) == 1 {
				slice_index, err := strconv.Atoi(pair[0])
				if err != nil {
					return Null{}, errors.New("slice index \"" + pair[0] + "\" is invalid ")
				}
				switch o := obj.(type) {
				case string:
					return o[slice_index], nil
				case []byte:
					return o[slice_index], nil
				case []int:
					return o[slice_index], nil
				case []string:
					return o[slice_index], nil
				case []map[string]interface{}:
					if len(remainder) > 0 {
						return Extract(remainder, o[slice_index])
					}
					return o[slice_index], nil
				case []map[string]string:
					if len(remainder) > 0 {
						return Extract(remainder, o[slice_index])
					}
					return o[slice_index], nil
				}
				return Null{}, errors.New("object \"" + fmt.Sprint(obj) + "\" (" + reflect.TypeOf(obj).String() + ") is not a slice.")
			}
			return Null{}, errors.New("slice range \"" + (path[1:slice_end_index]) + "\" is invalid ")
		} else {
			key := path[0:slice_start_index]
			remainder := path[slice_start_index:len(path)]
			switch o := obj.(type) {
			case Context:
				if value, ok := o[key]; ok {
					if value == nil {
						return Null{}, errors.New("value " + key + " is null.")
					}
					return Extract(remainder, value)
				}
				return Null{}, errors.New("value " + key + " is null.")
			case map[string]interface{}:
				if value, ok := o[key]; ok {
					if value == nil {
						return Null{}, errors.New("value " + key + " is null.")
					}
					return Extract(remainder, value)
				}
				return Null{}, errors.New("value " + key + " is null.")
			}
			return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())
		}
	}

	switch obj.(type) {
	case Context:
		if value, ok := obj.(Context)[path]; ok {
			if value == nil {
				return Null{}, nil
			}
			return value, nil
		}
		return Null{}, nil
	case map[string]interface{}:
		if value, ok := obj.(map[string]interface{})[path]; ok {
			if value == nil {
				return Null{}, nil
			}
			return value, nil
		}
		return Null{}, nil
	case map[string]string:
		if value, ok := obj.(map[string]string)[path]; ok {
			return value, nil
		}
		return Null{}, nil
	}

	return Null{}, errors.New("object is invalid type " + reflect.TypeOf(obj).String())

}
