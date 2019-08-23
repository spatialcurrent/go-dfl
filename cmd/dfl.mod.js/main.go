// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// dfl.mod.js is the package for go-dfl that is built as a JavaScript module.
// In modern JavaScript, the module can be imported using destructuring assignment.
// The functions are defined in the Exports variable in the dfljs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load functions into current scope
//	const { parse } = require('./dist/dfl.global.min.js);
//
//	// Parse an expression into a node
//	var { node, err } = parse("@a + @b")
//
//	// Compile the node
//	node = node.compile()
//
//	// Evaluate the node
//	var { result, err } = node.evaluate({"a": 1, "b": 2})
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/go-dfl/pkg/dfljs/
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/Destructuring_assignment
package main

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/spatialcurrent/go-dfl/pkg/dfljs"
)

func main() {
	jsModuleExports := js.Module.Get("exports")
	for name, value := range dfljs.Exports {
		jsModuleExports.Set(name, value)
	}
}
