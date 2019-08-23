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

	"github.com/pkg/errors"
)

// ParseCompileEvaluateInt parses the expression, compiles the node, evaluates on the given context, and returns a result of type int if valid, otherwise returns and error.
func ParseCompileEvaluateInt(exp string, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, int, error) {

	vars, value, err := ParseCompileEvaluate(exp, vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	switch value.(type) {
	case int:
		return vars, value.(int), nil
	}

	return vars, 0, errors.New("ParseCompileEvaluateInt evaluation returns invalid type " + fmt.Sprint(reflect.TypeOf(value)) + "")

}
