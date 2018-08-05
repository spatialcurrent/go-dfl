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
	"reflect"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// FunctionMap is a map of functions by string that are reference by name in the Function Node.
type FunctionMap map[string]func(Context, []interface{}) (interface{}, error)

func NewFuntionMapWithDefaults() FunctionMap {
	funcs := FunctionMap{}

	funcs["map"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 2 {
			return 0, errors.New("Invalid number of arguments to map.")
		}

		switch key := args[1].(type) {
		case string:
			switch a := args[0].(type) {
			case []map[string]interface{}:
				values := make([]interface{}, 0, len(a))
				for _, value := range a {
					values = append(values, value[key])
				}
				return values, nil
			case []map[string]string:
				values := make([]string, 0, len(a))
				for _, value := range a {
					values = append(values, value[key])
				}
				return values, nil
			}

			return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
		}

		return 0, errors.New("Invalid key for map function")

	}

	funcs["split"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 2 {
			return 0, errors.New("Invalid number of arguments to split.")
		}
		return strings.Split(fmt.Sprint(args[0]), fmt.Sprint(args[1])), nil
	}

	funcs["trim"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to split.")
		}

		switch a := args[0].(type) {
		case string:
			return strings.TrimSpace(a), nil
		case []byte:
			return []byte(strings.TrimSpace(string(a))), nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["len"] = func(ctx Context, args []interface{}) (interface{}, error) {
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

	funcs["bytes"] = func(ctx Context, args []interface{}) (interface{}, error) {
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

	funcs["int32"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to int32.")
		}
		switch a := args[0].(type) {
		case int:
			return int32(a), nil
		case int32:
			return a, nil
		case int64:
			return int32(a), nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["big"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to big.")
		}
		switch a := args[0].(type) {
		case int:
			buf := new(bytes.Buffer)
			err := binary.Write(buf, binary.BigEndian, int64(a))
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

	funcs["repeat"] = func(ctx Context, args []interface{}) (interface{}, error) {
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

	funcs["string"] = func(ctx Context, args []interface{}) (interface{}, error) {
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
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["min"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) < 1 {
			return 0, errors.New("Invalid number of arguments to len.")
		}

		return Min(TryConvertArray(args))
	}

	funcs["max"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) < 1 {
			return 0, errors.New("Invalid number of arguments to len.")
		}

		return Max(TryConvertArray(args))
	}

	funcs["lower"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}
		switch a := args[0].(type) {
		case string:
			return strings.ToLower(a), nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["upper"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}
		switch a := args[0].(type) {
		case string:
			return strings.ToUpper(a), nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["first"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}

		switch a := args[0].(type) {
		case string:
			return a[0], nil
		case []byte:
			return a[0], nil
		case []int:
			return a[0], nil
		case []float64:
			return a[0], nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["last"] = func(ctx Context, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}

		switch a := args[0].(type) {
		case string:
			return a[len(a)-1], nil
		case []byte:
			return a[len(a)-1], nil
		case []int:
			return a[len(a)-1], nil
		case []float64:
			return a[len(a)-1], nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	return funcs
}
