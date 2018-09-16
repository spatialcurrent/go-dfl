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

// ParseCompileEvaluateString parses the expression, compiles the node, evaluates on the given context, and returns a result of kind string if valid, otherwise returns and error.
func ParseCompileEvaluateString(exp string, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, string, error) {

	vars, value, err := ParseCompileEvaluate(exp, vars, ctx, funcs, quotes)
	if err != nil {
		return vars, "", err
	}

	switch value.(type) {
	case string:
		return vars, value.(string), nil
	}

	return vars, "", errors.New("ParseCompileEvaluateString evaluation returns invalid type " + fmt.Sprint(reflect.TypeOf(value)) + "")

}
