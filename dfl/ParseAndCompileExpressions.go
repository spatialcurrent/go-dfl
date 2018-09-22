// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
)

// ParseAndCompileExpressions is a usability function to parse and compile multiple expressions.
func ParseAndCompileExpressions(expressions map[string]string) (map[string]Node, error) {
	nodes := map[string]Node{}
	for k, exp := range expressions {
		node, _, err := Parse(exp)
		if err != nil {
			return nodes, errors.Wrap(err, "error parsing expression "+exp)
		}
		nodes[k] = node.Compile()
	}
	return nodes, nil
}
