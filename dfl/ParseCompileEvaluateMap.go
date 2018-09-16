// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

// ParseCompileEvaluateMap parses the expression, compiles the node, evaluates on the given context, and returns a result of kind map if valid, otherwise returns and error.
func ParseCompileEvaluateMap(exp string, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, value, err := ParseCompileEvaluate(exp, vars, ctx, funcs, quotes)
	if err != nil {
		return vars, map[string]interface{}{}, err
	}

	if reflect.TypeOf(value).Kind() != reflect.Map {
		return vars, value, errors.New("ParseCompileEvaluateString evaluation returns invalid type " + fmt.Sprint(reflect.TypeOf(value)) + "")
	}

	return vars, value, nil
}
