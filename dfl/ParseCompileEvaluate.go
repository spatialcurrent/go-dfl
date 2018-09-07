// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

func ParseCompileEvaluate(exp string, vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	node, err := Parse(exp)
	if err != nil {
		return vars, Null{}, err
	}
	return node.Compile().Evaluate(vars, ctx, funcs, quotes)
}
