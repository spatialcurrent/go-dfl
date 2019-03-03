// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
)

// MustParseCompile parses the input expression and compiles the DFL node.  Panics if any error.
func MustParseCompile(expression string) Node {
	node, remainder, err := Parse(expression)
	if err != nil {
		panic(err)
	}
	if len(remainder) > 0 {
		panic(errors.New("MustParseCompile cannot have a remainder: " + remainder))
	}
	return node.Compile()
}
