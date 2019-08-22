// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// dfl.global.js is the package for go-dfl that adds DFL functions to the global scope under the "dfl" variable.
//
// In Node, depending on where require is called and the build system used, the functions may need to be required at the top of each module file.
// In a web browser, gss can be made available to the entire web page.
// The functions are defined in the Exports variable in the gssjs package.
//
// Usage
//	// Below is a simple set of examples of how to use this package in a JavaScript application.
//
//	// load functions into global scope
//	// require('./dist/dfl.global.min.js);
//
//	// Parse an expression into a node
//	var { node, err } = dfl.parse("@a + @b")
//
//	// Compile the node
//	node = node.compile()
//
//	// Evaluate the node
//	var { result, err } = node.evaluate({"a": 1, "b": 2})
//
// References
//	- https://godoc.org/pkg/github.com/spatialcurrent/go-dfl/pkg/dfljs/
//	- https://nodejs.org/api/globals.html#globals_global_objects
//	- https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects
package main

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/spatialcurrent/go-dfl/pkg/dfljs"
)

func main() {
	js.Global.Set("dfl", dfljs.Exports)
}
