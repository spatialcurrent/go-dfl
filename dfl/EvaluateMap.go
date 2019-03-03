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
)

import (
	"github.com/pkg/errors"
)

// EvaluateMap returns the map value of a node given a context.  If the result is not a map, then returns an error.
func EvaluateMap(n Node, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, result, err := n.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, "", errors.Wrap(err, "Error evaluating expression "+n.Dfl(quotes, false, 0))
	}

	if reflect.TypeOf(result).Kind() != reflect.Map {
		return vars, "", errors.New("evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of a map")
	}

	return vars, result, nil
}
