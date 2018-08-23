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
// In you html document, the simplest workflow is to add dfl as a script and call dfl.EvaluateBool(expression, {"a": 1});
// For performance reasons, you can compile your expression as below:
//	var exp = "@pop + 10";
//	var root = dfl.Parse(exp).Compile();
//	var result = root.Evaluate({"pop": 10})
//
package main

import (
	"github.com/spatialcurrent/go-dfl/dfl"
)

import (
	"github.com/gopherjs/gopherjs/js"
	"honnef.co/go/js/console"
)

type Node struct {
	Node        dfl.Node
	FunctionMap dfl.FunctionMap
}

func (n Node) Compile() *js.Object {
	return js.MakeWrapper(Node{
		Node:        n.Node.Compile(),
		FunctionMap: dfl.NewFuntionMapWithDefaults(),
	})
}

func (n Node) Evaluate(options *js.Object) interface{} {

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	result, err := n.Node.Evaluate(ctx, n.FunctionMap)
	if err != nil {
		console.Log(err.Error())
		return false
	}
	return result
}

func main() {
	js.Global.Set("dfl", map[string]interface{}{
		"version":        dfl.VERSION,
		"Parse":          Parse,
		"EvaluateBool":   EvaluateBool,
		"EvaluateInt":    EvaluateInt,
		"EvaluateFloat":  EvaluateFloat64,
		"EvaluateString": EvaluateString,
	})
}

func Parse(s string) *js.Object {
	root, err := dfl.Parse(s)
	if err != nil {
		console.Log(err.Error())
		return js.MakeWrapper(Node{Node: nil})
	}
	return js.MakeWrapper(Node{Node: root})
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

	result, err := dfl.EvaluateBool(root, ctx, dfl.NewFuntionMapWithDefaults())
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

	result, err := dfl.EvaluateInt(root, ctx, dfl.NewFuntionMapWithDefaults())
	if err != nil {
		console.Log(err.Error())
		return 0
	}

	return result
}

func EvaluateFloat64(s string, options *js.Object) float64 {
	root, err := dfl.Parse(s)
	if err != nil {
		console.Log(err.Error())
		return 0.0
	}

	root = root.Compile()

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	result, err := dfl.EvaluateFloat64(root, ctx, dfl.NewFuntionMapWithDefaults())
	if err != nil {
		console.Log(err.Error())
		return 0.0
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

	result, err := dfl.EvaluateString(root, ctx, dfl.NewFuntionMapWithDefaults())
	if err != nil {
		console.Log(err.Error())
		return ""
	}

	return result
}
