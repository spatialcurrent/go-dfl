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

// EvaluateFloat64 returns the float64 value of a node given a context.  If the result is not a float64, then returns an error.
func EvaluateFloat64(n Node, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, float64, error) {
	vars, result, err := n.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0.0, errors.Wrap(err, "Error evaluating expression "+n.Dfl(quotes, false, 0))
	}

	switch result.(type) {
	case int:
		return vars, float64(result.(int)), nil
	case float64:
		return vars, result.(float64), nil
	}

	return vars, 0.0, errors.New("Evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of int")
}
