// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfljs

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
	"honnef.co/go/js/console"

	"github.com/spatialcurrent/go-dfl/dfl"
)

// EvaluateBool provides a simple function that parses, compiles, and executes an expression against a context object and returns true or false.
func EvaluateBool(s string, options *js.Object) bool {
	root, err := dfl.ParseCompile(s)
	if err != nil {
		console.Error(errors.Wrap(err, "error parsing expression in EvaluateBool").Error())
		return false
	}

	vars := map[string]interface{}{}

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	_, result, err := dfl.EvaluateBool(root, vars, ctx, dfl.NewFuntionMapWithDefaults(), DefaultQuotes[1:])
	if err != nil {
		console.Error(errors.Wrap(err, "error evaluating a node in EvaluateBool").Error())
		return false
	}

	return result
}
