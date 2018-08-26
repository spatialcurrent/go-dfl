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

// EvaluateString returns the string value of a node given a context.  If the result is not a string, then returns an error.
func EvaluateString(n Node, ctx interface{}, funcs FunctionMap, quotes []string) (string, error) {
	result, err := n.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return "", errors.Wrap(err, "Error evaluating expression "+n.Dfl(quotes, false))
	}

	switch result.(type) {
	case string:
		return result.(string), nil
	}

	return "", errors.New("Evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of string")
}
