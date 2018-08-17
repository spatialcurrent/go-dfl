// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// FunctionMap is a map of functions by string that are reference by name in the Function Node.
type FunctionMap map[string]func(interface{}, []interface{}) (interface{}, error)

func NewFuntionMapWithDefaults() FunctionMap {
	funcs := FunctionMap{}

	funcs["map"] = mapArray
	funcs["len"] = getLength
	funcs["bytes"] = convertToBytes
	funcs["int16"] = convertToInt16
	funcs["int32"] = convertToInt32
	funcs["big"] = convertToBigEndian
	funcs["little"] = convertToLittleEndian
	funcs["repeat"] = repeat
	funcs["string"] = convertToString
	funcs["split"] = splitString
	funcs["prefix"] = prefix
	funcs["suffix"] = suffix
	funcs["trim"] = trimString
	funcs["ltrim"] = trimStringLeft
	funcs["rtrim"] = trimStringRight

	funcs["min"] = func(ctx interface{}, args []interface{}) (interface{}, error) {
		if len(args) < 1 {
			return 0, errors.New("Invalid number of arguments to len.")
		}

		return Min(TryConvertArray(args))
	}

	funcs["max"] = func(ctx interface{}, args []interface{}) (interface{}, error) {
		if len(args) < 1 {
			return 0, errors.New("Invalid number of arguments to len.")
		}

		return Max(TryConvertArray(args))
	}

	funcs["lower"] = func(ctx interface{}, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}
		switch a := args[0].(type) {
		case string:
			return strings.ToLower(a), nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["upper"] = func(ctx interface{}, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}
		switch a := args[0].(type) {
		case string:
			return strings.ToUpper(a), nil
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["first"] = func(ctx interface{}, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}

		return First(args[0])
	}

	funcs["last"] = func(ctx interface{}, args []interface{}) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}

		return Last(args[0])
	}

	return funcs
}
