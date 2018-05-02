// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
	"unicode"
)

import (
	"github.com/pkg/errors"
)

// ParseAttribute parses an Attribute Node from an input string
// If parameter "in" is gramatically a child node, then return the parent node.
// All attribute references must start with an "@" character and have no spaces.
// For example, @amenity, @shop, @population, etc.
// Given those rules the remainder, if any, if simply parsed from the input strings
// Examples:
//	node, err := ParseAttribute("@amenity in [bar, restaurant]")
func ParseAttribute(in string) (Node, error) {

	end := strings.Index(strings.TrimLeftFunc(in, unicode.IsSpace), " ")
	if end == -1 {
		return &Attribute{Name: strings.TrimSpace(in)[1:]}, nil
	}

	if len(strings.TrimSpace(in[end:])) == 0 {
		return &Attribute{Name: in[1:end]}, nil
	}

	left := &Attribute{Name: in[1:end]}
	root, err := Parse(in[end:])
	if err != nil {
		return root, err
	}
	switch root.(type) {
	case *And:
		root.(*And).Left = left
	case *Or:
		root.(*Or).Left = left
	case *Xor:
		root.(*Xor).Left = left
	case *In:
		root.(*In).Left = left
	case *Like:
		root.(*Like).Left = left
	case *ILike:
		root.(*ILike).Left = left
	case *LessThan:
		root.(*LessThan).Left = left
	case *LessThanOrEqual:
		root.(*LessThanOrEqual).Left = left
	case *GreaterThan:
		root.(*GreaterThan).Left = left
	case *GreaterThanOrEqual:
		root.(*GreaterThanOrEqual).Left = left
	case *Equal:
		root.(*Equal).Left = left
	case *NotEqual:
		root.(*NotEqual).Left = left
	case *Add:
		root.(*Add).Left = left
	case *Subtract:
		root.(*Subtract).Left = left
	case *Before:
		root.(*Before).Left = left
	case *After:
		root.(*After).Left = left
	default:
		return root, errors.New("Invalid expression syntax for " + in + ".  Root is not a binary operator")
	}
	return root, nil

}
