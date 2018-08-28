// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

// BinaryOperator is a DFL Node that represents the binary operator of a left value and right value.
// This struct functions as an embedded struct for many comparator operations.
type BinaryOperator struct {
	Left  Node
	Right Node
}

func (bo BinaryOperator) Dfl(operator string, quotes []string, pretty bool, tabs int) string {
	if pretty {
		switch bo.Left.(type) {
		case *Literal:
			switch bo.Left.(*Literal).Value.(type) {
			case string, int, []byte, Null:
				return strings.Repeat("  ", tabs) + "(" + bo.Left.Dfl(quotes, false, tabs) + " " + operator + " " + bo.Right.Dfl(quotes, false, tabs) + ")"
			}
		}
		switch bo.Right.(type) {
		case *Literal:
			switch bo.Right.(*Literal).Value.(type) {
			case string, int, []byte, Null:
				return strings.Repeat("  ", tabs) + "(" + bo.Left.Dfl(quotes, false, tabs) + " " + operator + " " + bo.Right.Dfl(quotes, false, tabs) + ")"
			}
		}
		return strings.Repeat("  ", tabs) + "(\n" + bo.Left.Dfl(quotes, pretty, tabs+1) + " " + operator + " " + "\n" + bo.Right.Dfl(quotes, pretty, tabs+1) + "\n" + strings.Repeat("  ", tabs) + ")"
	}
	return "(" + bo.Left.Dfl(quotes, pretty, tabs) + " " + operator + " " + bo.Right.Dfl(quotes, pretty, tabs) + ")"
}

// EvaluateLeftAndRight evaluates the value of the left node and right node given a context map (ctx) and function map (funcs).
// Returns a 3 value tuple of left value, right value, and error.
func (bo BinaryOperator) EvaluateLeftAndRight(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, interface{}, error) {
	lv, err := bo.Left.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return false, false, err
	}
	rv, err := bo.Right.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return false, false, err
	}
	return lv, rv, nil
}

// Attributes returns a slice of all attributes used in the evluation of this node, including a children nodes.
// Attributes de-duplicates values from the left node and right node using a set.
func (bo BinaryOperator) Attributes() []string {
	set := make(map[string]struct{})
	for _, x := range bo.Left.Attributes() {
		set[x] = struct{}{}
	}
	for _, x := range bo.Right.Attributes() {
		set[x] = struct{}{}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
