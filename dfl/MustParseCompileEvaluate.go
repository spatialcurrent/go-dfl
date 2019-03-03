// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// MustParseCompileEvaluate parses the expression, compiles the node, and evaluates on the given context.  Panics if any error.
func MustParseCompileEvaluate(exp string, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}) {
	vars, result, err := ParseCompileEvaluate(exp, vars, ctx, funcs, quotes)
	if err != nil {
		panic(err)
	}

	return vars, result
}
