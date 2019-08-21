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

// EvaluateString returns the string value of a node given a context.  If the result is not a string, then returns an error.
func EvaluateString(n Node, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, string, error) {
	vars, result, err := n.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, "", errors.Wrap(err, "Error evaluating expression "+n.Dfl(quotes, false, 0))
	}

	switch result.(type) {
	case string:
		return vars, result.(string), nil
	}

	return vars, "", errors.New("Evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of string")
}
