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
		return root, errors.New("Invalid expression syntax for " + s + ".  Root is not a binary operator")
	}

	return root, nil
}
