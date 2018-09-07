// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
	"regexp"
	"strings"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-adaptive-functions/af"
)

// FunctionMap is a map of functions by string that are reference by name in the Function Node.
type FunctionMap map[string]func(FunctionMap, map[string]interface{}, interface{}, []interface{}, []string) (interface{}, error)

func NewFuntionMapWithDefaults() FunctionMap {
	funcs := FunctionMap{}

	for _, fn := range af.Functions {
		for _, alias := range fn.Aliases {
			funcs[alias] = func(fn af.Function) func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
				return func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
					if err := fn.Validate(args); err != nil {
						return Null{}, errors.Wrap(err, "Invalid arguments")
					}
					return fn.Run(args)
				}
			}(fn)
		}
	}

	funcs["filter"] = filterArray
	funcs["map"] = mapArray
	funcs["sort"] = sortArray
	funcs["dict"] = toDict
	funcs["hist"] = histArray
	funcs["prefix"] = prefix
	funcs["suffix"] = suffix
	funcs["trim"] = trimString
	funcs["ltrim"] = trimStringLeft
	funcs["rtrim"] = trimStringRight

	funcs["slugify"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) < 2 {
			return 0, errors.New("Invalid number of arguments to slugify.")
		}

		switch s := args[0].(type) {
		case string:
			switch replacement := args[1].(type) {
			case string:
				reg, err := regexp.Compile("[^a-zA-Z0-9]+")
				if err != nil {
					return Null{}, errors.Wrap(err, "Invalid regular expression ")
				}
				return reg.ReplaceAllString(strings.ToLower(s), replacement), nil
			}
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["min"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) < 1 {
			return 0, errors.New("Invalid number of arguments to len.")
		}

		return Min(TryConvertArray(args))
	}

	funcs["max"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) < 1 {
			return 0, errors.New("Invalid number of arguments to len.")
		}

		return Max(TryConvertArray(args))
	}

	funcs["first"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}

		return First(args[0])
	}

	funcs["last"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to upper.")
		}

		return Last(args[0])
	}

	return funcs
}
