// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package dfljs includes functions for the JavaScript build of go-dfl.
//
package dfljs

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

// DefaultQuotes is the default quotes for JavaScript.
var DefaultQuotes = []string{"\"", "'", "`"}

func isArray(object *js.Object) bool {
	return js.Global.Get("Array").Call("isArray", object).Bool()
}

func toArray(object *js.Object) []interface{} {
	arr := make([]interface{}, 0, object.Length())
	for i := 0; i < object.Length(); i++ {
		arr = append(arr, parseObject(object.Index(i)))
	}
	return arr
}

func parseObject(object *js.Object) interface{} {
	if isArray(object) {
		return toArray(object)
	}
	return object.Interface()
}

// Node is a struct that wraps a dfl.Node and dfl.FunctionMap to provide an api to dfl.js.
type Node struct {
	Node        dfl.Node
	FunctionMap dfl.FunctionMap
}

func (n Node) Pretty() string {
	return n.Node.Dfl(DefaultQuotes[1:], true, 0)
}

func (n Node) Compile() *js.Object {
	return js.MakeWrapper(Node{
		Node:        n.Node.Compile(),
		FunctionMap: dfl.NewFuntionMapWithDefaults(),
	})
}

func (n Node) Evaluate(options *js.Object) map[string]interface{} {
	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}
	_, result, err := n.Node.Evaluate(map[string]interface{}{}, ctx, n.FunctionMap, DefaultQuotes[1:])
	if err != nil {
		return map[string]interface{}{"err": errors.Wrap(err, "error evaluating").Error()}
	}
	return map[string]interface{}{"result": result, "err": err}
}

func exports() map[string]interface{} {
	m := map[string]interface{}{}

	m["fmt"] = func(s string) map[string]interface{} {
		node, remainder, err := dfl.Parse(s)
		if err != nil {
			return map[string]interface{}{"err": errors.Wrap(err, "error parsing an expression").Error()}
		}
		expression := node.Dfl(DefaultQuotes, false, 0)
		return map[string]interface{}{"expression": expression, "remainder": remainder, "err": err}
	}

	m["parse"] = func(s string) map[string]interface{} {
		root, remainder, err := dfl.Parse(s)
		if err != nil {
			return map[string]interface{}{"err": errors.Wrap(err, "error parsing an expression").Error()}
		}
		node := js.MakeWrapper(Node{Node: root, FunctionMap: dfl.NewFuntionMapWithDefaults()})
		return map[string]interface{}{"node": node, "remainder": remainder, "err": err}
	}

	m["compile"] = func(node *js.Object) *js.Object {
		return node.Call("Compile")
	}

	m["evaluate"] = func(node *js.Object, ctx *js.Object) *js.Object {
		return node.Call("Evaluate", ctx)
	}

	return m
}

var Exports = exports()
