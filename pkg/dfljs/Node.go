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
	"github.com/spatialcurrent/go-dfl/dfl"
	"honnef.co/go/js/console"
)

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

func (n Node) Evaluate(options *js.Object) interface{} {

	ctx := map[string]interface{}{}
	for _, key := range js.Keys(options) {
		ctx[key] = options.Get(key).Interface()
	}

	_, result, err := n.Node.Evaluate(map[string]interface{}{}, ctx, n.FunctionMap, DefaultQuotes[1:])
	if err != nil {
		console.Error(errors.Wrap(err, "error evaluating").Error())
		return false
	}
	return result
}
