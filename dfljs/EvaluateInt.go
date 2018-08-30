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

// EvaluateInt provides a simple function that parses, compiles, and executes an expression against a context object and returns an integer.
func EvaluateInt(s string, options *js.Object) int {
	root, err := dfl.Parse(s)
	if err != nil {
		console.Error(errors.Wrap(err, "error parsing expression in EvaluateInt").Error())
		return 0
	}

	root = root.Compile()

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	result, err := dfl.EvaluateInt(root, ctx, dfl.NewFuntionMapWithDefaults(), DefaultQuotes[1:])
	if err != nil {
		console.Error(errors.Wrap(err, "error evaluating a node in EvaluateInt").Error())
		return 0
	}

	return result
}
