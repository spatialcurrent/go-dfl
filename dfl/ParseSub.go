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

// ParseSub is used to parse a sub-expression and the remainder, if any.
// A sub-expression is usually enclosed by parantheses.  The parantheses are removed before being passed to ParseSub.
// If parameter "in" is gramatically a child node, then return the parent node.
// For Example with an input string "(@cuisine like mexican) or (@name ilike %burrito%)",
//	node, err : ParseSub("@cuisine like mexican", "or (@name ilike %burrito%)")
//
func ParseSub(s string, remainder string) (Node, error) {

	if len(remainder) == 0 {
		return Parse(s)
	}

	var root Node
	left, err := Parse(s)
	if err != nil {
		return root, err
	}

	root, err = Parse(remainder)
	if err != nil {
		return root, err
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, errors.Wrap(err, "error attaching left for "+s)
	}

	return root, nil
}
