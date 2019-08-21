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

// EvaluateString provides a simple function that parses, compiles, and executes an expression against a context object and returns a string.
func EvaluateString(s string, options *js.Object) string {
	root, err := dfl.ParseCompile(s)
	if err != nil {
		console.Error(errors.Wrap(err, "error parsing expression in EvaluateString").Error())
		return ""
	}

	vars := map[string]interface{}{}

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	_, result, err := dfl.EvaluateString(root, vars, ctx, dfl.NewFuntionMapWithDefaults(), DefaultQuotes[1:])
	if err != nil {
		console.Error(errors.Wrap(err, "error evaluating a node in EvaluateString").Error())
		return ""
	}

	return result
}
