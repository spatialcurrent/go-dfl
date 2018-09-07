// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
	"strings"
	"unicode"
)

// ParseVariable parses a Variable Node from an input string
// If parameter "in" is gramatically a child node, then return the parent node.
// All variable references must start with an "$" character and have no spaces.
// For example, $amenity, $shop, $population, etc.
// Given those rules the remainder, if any, if simply parsed from the input strings
// Examples:
//	node, err := ParseAttribute("$amenities := [bar, restaurant]")
func ParseVariable(in string, remainder string) (Node, error) {

	if len(remainder) == 0 {
		return &Variable{Name: strings.TrimLeftFunc(in, unicode.IsSpace)[1:]}, nil
	}

	left := &Variable{Name: strings.TrimLeftFunc(in, unicode.IsSpace)[1:]}
	root, err := Parse(remainder)
	if err != nil {
		return root, errors.Wrap(err, "error parsing remainder < "+remainder+" >")
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, errors.Wrap(err, "error attaching left for variable "+in)
	}

	return root, nil

}
