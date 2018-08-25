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

// ParseArray parses an array of nodes.
// If parameter "in" is gramatically a child node, then return the parent node.
// DFL arrays can include Attribute or Literal Nodes.
// As all attribute references must start with an "@" character, parantheses are optional for literals except if a comma exists.
// Below are some example inputs
//
//	[bank, bureau_de_change, atm]
//	[1, 2, @target]
//	[Taco, Tacos, Burrito, Burritos, "Mexican Food", @example]
func ParseArray(in string, remainder string) (Node, error) {

	nodes, err := ParseList(in)
	if err != nil {
		return &Array{}, errors.Wrap(err, "error parsing array "+in)
	}

	if len(remainder) == 0 {
		return &Array{Nodes: nodes}, nil
	}

	left := &Array{Nodes: nodes}
	root, err := Parse(remainder)
	if err != nil {
		return root, err
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, errors.Wrap(err, "error attaching left for "+in)
	}

	return root, nil

}
