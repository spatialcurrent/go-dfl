// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
)

import (
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
func ParseLiteral(v interface{}, remainder string) (Node, error) {

	if len(remainder) == 0 {
		return &Literal{Value: v}, nil
	}

	left := &Literal{Value: v}
	root, err := Parse(remainder)
	if err != nil {
		return root, errors.Wrap(err, "error parsing remainder < "+remainder+" >")
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, errors.Wrap(err, "Could not attach left "+fmt.Sprint(v))
	}

	return root, nil
}
