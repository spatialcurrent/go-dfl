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

// EvaluateMap returns the map value of a node given a context.  If the result is not a map, then returns an error.
func EvaluateMap(n Node, ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {
	_, result, err := n.Evaluate(map[string]interface{}{}, ctx, funcs, quotes)
	if err != nil {
		return "", errors.Wrap(err, "Error evaluating expression "+n.Dfl(quotes, false, 0))
	}

	if reflect.TypeOf(result).Kind() != reflect.Map {
		return "", errors.New("evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of a map")
	}

	return result, nil
}
