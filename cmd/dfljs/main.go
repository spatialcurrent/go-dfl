// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// DFLJS is the Javascript version of DFL.
//
// Usage
//
// In you html document simply add dfl as a script and call dfl.EvaluateBool(expression, {"a": 1});
//
package main

import (
	"github.com/spatialcurrent/go-dfl/dfl"
)

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/console"
)

var GO_DFL_VERSION = "0.0.2"

func main() {
	js.Global.Set("dfl", map[string]interface{}{
		"Version":        GO_DFL_VERSION,
		"EvaluateBool":   EvaluateBool,
		"EvaluateInt":    EvaluateInt,
		"EvaluateString": EvaluateString,
	})
}

func EvaluateBool(s string, options *js.Object) bool {
	root, err := dfl.Parse(s)
	if err != nil {
		console.Log(err.Error())
		return false
	}

	root = root.Compile()

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	result, err := dfl.EvaluateBool(root, ctx, dfl.FunctionMap{})
	if err != nil {
		console.Log(err.Error())
		return false
	}

	return result
}

func EvaluateInt(s string, options *js.Object) int {
	root, err := dfl.Parse(s)
	if err != nil {
		console.Log(err.Error())
		return 0
	}

	root = root.Compile()

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	result, err := dfl.EvaluateInt(root, ctx, dfl.FunctionMap{})
	if err != nil {
		console.Log(err.Error())
		return 0
	}

	return result
}

func EvaluateString(s string, options *js.Object) string {
	root, err := dfl.Parse(s)
	if err != nil {
		console.Log(err.Error())
		return ""
	}

	root = root.Compile()

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	result, err := dfl.EvaluateString(root, ctx, dfl.FunctionMap{})
	if err != nil {
		console.Log(err.Error())
		return ""
	}

	return result
}
