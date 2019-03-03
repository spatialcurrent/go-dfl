// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
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
	"github.com/gopherjs/gopherjs/js"
	"github.com/spatialcurrent/go-dfl/dfl"
	"github.com/spatialcurrent/go-dfl/dfljs"
)

func main() {
	js.Global.Set("dfl", map[string]interface{}{
		"version":        dfl.Version,
		"Parse":          dfljs.Parse,
		"EvaluateBool":   dfljs.EvaluateBool,
		"EvaluateInt":    dfljs.EvaluateInt,
		"EvaluateFloat":  dfljs.EvaluateFloat64,
		"EvaluateMap":    dfljs.EvaluateMap,
		"EvaluateString": dfljs.EvaluateString,
	})
}
