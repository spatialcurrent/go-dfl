// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// ParseCompileEvaluate parses the expression, compiles the node, and evaluates on the given context.
func ParseCompileEvaluate(exp string, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	node, err := ParseCompile(exp)
	if err != nil {
		return vars, Null{}, err
	}

	return node.Evaluate(vars, ctx, funcs, quotes)
}
