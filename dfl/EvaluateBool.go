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

// EvaluateBool returns the boolean value of a node given a context.  If the result is not a bool, then returns an error.
func EvaluateBool(n Node, ctx interface{}, funcs FunctionMap) (bool, error) {
	result, err := n.Evaluate(ctx, funcs)
	if err != nil {
		return false, errors.Wrap(err, "Error evaluating expression "+n.Dfl())
	}

	switch result.(type) {
	case bool:
		return result.(bool), nil
	}

	return false, errors.New("Evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of bool")
}
