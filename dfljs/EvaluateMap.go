// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfljs

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-dfl/dfl"
	"honnef.co/go/js/console"
)

// EvaluateMap provides a simple function that parses, compiles, and executes an expression against a context object and returns a map.
func EvaluateMap(s string, options *js.Object) interface{} {
	root, err := dfl.Parse(s)
	if err != nil {
		console.Error(errors.Wrap(err, "error parsing expression in EvaluateMap").Error())
		return ""
	}

	root = root.Compile()

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	result, err := dfl.EvaluateMap(root, ctx, dfl.NewFuntionMapWithDefaults(), DefaultQuotes[1:])
	if err != nil {
		console.Error(errors.Wrap(err, "error evaluating a node in EvaluateMap").Error())
		return ""
	}

	return result
}
