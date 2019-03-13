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

// ParseSub is used to parse a sub-expression and the remainder, if any.
// A sub-expression is usually enclosed by parantheses.  The parantheses are removed before being passed to ParseSub.
// If parameter "in" is gramatically a child node, then return the parent node.
// For Example with an input string "(@cuisine like mexican) or (@name ilike %burrito%)",
//	node, err : ParseSub("@cuisine like mexican", "or (@name ilike %burrito%)")
//
func ParseSub(s string, remainder string) (Node, string, error) {

	// If there is no remainder
	if len(remainder) == 0 {
		return Parse(s)
	}

	// If the remainder starts with a break
	if len(remainder) >= 2 && remainder[0] == ':' && remainder[1] == ' ' {
		n, subRemainder, err := Parse(s)
		if err != nil {
			return n, remainder, err
		}
		if len(subRemainder) > 0 {
			return n, remainder, errors.New("interior sub should not have a remainder: " + subRemainder)
		}

		return n, remainder, err
	}

	var root Node
	left, _, err := Parse(s)
	if err != nil {
		return root, "", err
	}

	root, remainder, err = Parse(remainder)
	if err != nil {
		return root, remainder, err
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, remainder, errors.Wrap(err, "error attaching left for "+s)
	}

	return root, remainder, nil

}
