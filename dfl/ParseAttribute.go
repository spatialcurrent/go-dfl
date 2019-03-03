// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
	"strings"
	"unicode"
)

// ParseAttribute parses an Attribute Node from an input string
// If parameter "in" is gramatically a child node, then return the parent node.
// All attribute references must start with an "@" character and have no spaces.
// For example, @amenity, @shop, @population, etc.
// Given those rules the remainder, if any, if simply parsed from the input strings
// Examples:
//	node, err := ParseAttribute("@amenity in [bar, restaurant]")
func ParseAttribute(in string, remainder string) (Node, string, error) {

	if len(remainder) == 0 || (len(remainder) >= 2 && remainder[0] == ':' && remainder[1] != '=') {
		return &Attribute{Name: strings.TrimLeftFunc(in, unicode.IsSpace)[1:]}, remainder, nil
	}

	left := &Attribute{Name: strings.TrimLeftFunc(in, unicode.IsSpace)[1:]}
	root, remainder, err := Parse(remainder)
	if err != nil {
		return root, remainder, errors.Wrap(err, "error parsing remainder < "+remainder+" >")
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, remainder, errors.Wrap(err, "error attaching left for attribute "+in)
	}

	return root, remainder, nil

}
