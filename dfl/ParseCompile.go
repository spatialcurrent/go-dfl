// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// ParseCompile parses the input expression and compiles the DFL node.
func ParseCompile(expression string) (Node, error) {
	node, _, err := Parse(expression)
	if err != nil {
		return node, err
	}
	return node.Compile(), nil
}
