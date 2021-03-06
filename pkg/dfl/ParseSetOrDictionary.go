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

// ParseSetOrDictionary parses a Set or Dictionary Node and recursively any remainder.
// If parameter "in" is gramatically a child node, then return the parent node.
// DFL sets/dictionaries can include Attribute or Literal Nodes.
// As all attribute references must start with an "@" character, parentheses are optional for literals except if a comma exists.
// Below are some example inputs
//
//	{bank, bureau_de_change, atm}
//	{1, 2, @target}
//	{Taco, Tacos, Burrito, Burritos, "Mexican Food", @example}
//	{amenity: bank}
func ParseSetOrDictionary(in string, remainder string) (Node, string, error) {

	isSet, list, kv, err := ParseListOrKeyValue(in)
	if err != nil {
		return &Set{}, remainder, errors.Wrap(err, "error parsing set "+in)
	}

	if len(remainder) == 0 || (len(remainder) >= 2 && remainder[0] == ':' && remainder[1] != '=') {
		if isSet {
			return &Set{Nodes: list}, remainder, nil
		}
		return &Dictionary{Nodes: kv}, remainder, nil
	}

	root, remainder, err := Parse(remainder)
	if err != nil {
		return root, remainder, err
	}

	if isSet {
		err = AttachLeft(root, &Set{Nodes: list})
	} else {
		err = AttachLeft(root, &Dictionary{Nodes: kv})
	}

	if err != nil {
		return root, remainder, errors.Wrap(err, "error attaching left for "+in)
	}

	return root, remainder, nil

}
