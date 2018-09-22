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
)

import (
	"github.com/pkg/errors"
)

// EvaluateArray returns the array/slice value of a node given a context.  If the result is not an array or slice, then returns an error.
func EvaluateArray(n Node, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, result, err := n.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, "", errors.Wrap(err, "Error evaluating expression "+n.Dfl(quotes, false, 0))
	}

	if t := reflect.TypeOf(result); t.Kind() != reflect.Array && t.Kind() != reflect.Slice {
		return vars, "", errors.New("evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of an array or slice")
	}

	return vars, result, nil
}
