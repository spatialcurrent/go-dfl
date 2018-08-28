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

// EvaluateFloat64 returns the float64 value of a node given a context.  If the result is not a float64, then returns an error.
func EvaluateFloat64(n Node, ctx interface{}, funcs FunctionMap, quotes []string) (float64, error) {
	result, err := n.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return 0.0, errors.Wrap(err, "Error evaluating expression "+n.Dfl(quotes, false, 0))
	}

	switch result.(type) {
	case int:
		return float64(result.(int)), nil
	case float64:
		return result.(float64), nil
	}

	return 0.0, errors.New("Evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of int")
}
