// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"

	"github.com/pkg/errors"
)

// ParseLiteral wraps parameter v in a Literal Node and parses a remainder, if any.
// ParseLiteral does not additional parsing of parameter v.
// TryConvertString is used to parse an int, float64, or time from a string representation.
// If parameter "in" is gramatically a child node, then return the parent node.
// For example, @amenity, @shop, @population, etc.
// Given those rules the remainder, if any, if simply parsed from the input strings
// Examples:
//	node, err := ParseLiteral("brewery")
func ParseLiteral(v interface{}, remainder string) (Node, string, error) {

	if len(remainder) == 0 || (len(remainder) >= 2 && remainder[0] == ':' && remainder[1] != '=') {
		return &Literal{Value: v}, remainder, nil
	}

	left := &Literal{Value: v}
	root, remainder, err := Parse(remainder)
	if err != nil {
		return root, remainder, errors.Wrap(err, "error parsing remainder < "+remainder+" >")
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, remainder, errors.Wrap(err, "could not attach left "+fmt.Sprint(v)+" to root "+fmt.Sprint(root))
	}

	return root, remainder, nil
}
