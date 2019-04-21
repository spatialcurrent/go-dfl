// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/spatialcurrent/go-dfl/dfl/builder"
)

// BinaryOperator is a DFL Node that represents the binary operator of a left value and right value.
// This struct functions as an embedded struct for many comparator operations.
type BinaryOperator struct {
	Left  Node
	Right Node
}

func (bo BinaryOperator) Builder(operator string, quotes []string, tabs int) builder.Builder {
	return builder.New(quotes, tabs).Left(bo.Left).Op(operator).Right(bo.Right)
}

func (bo BinaryOperator) Dfl(operator string, quotes []string, pretty bool, tabs int) string {
	b := bo.Builder(operator, quotes, tabs)
	if pretty {
		b = b.Indent(tabs)
		switch bo.Left.(type) {
		case *Literal:
			switch bo.Left.(*Literal).Value.(type) {
			case string, int, []byte, Null:
				return b.Dfl()
			}
		}
		switch bo.Right.(type) {
		case *Literal:
			switch bo.Right.(*Literal).Value.(type) {
			case string, int, []byte, Null:
				return b.Dfl()
			}
		}
		return b.Pretty(true).Tabs(tabs).Dfl()
	}
	return b.Dfl()
}

func (bo BinaryOperator) Sql(operator string, pretty bool, tabs int) string {
	return builder.New([]string{}, tabs).Left(bo.Left).Op(operator).Right(bo.Right).Sql()
}

func (bo BinaryOperator) Map(operator string, left Node, right Node) map[string]interface{} {
	return map[string]interface{}{
		"op":    operator,
		"left":  left.Map(),
		"right": right.Map(),
	}
}

// EvaluateLeftAndRight evaluates the value of the left node and right node given a context map (ctx) and function map (funcs).
// Returns a 3 value tuple of left value, right value, and error.
func (bo BinaryOperator) EvaluateLeftAndRight(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, interface{}, error) {

	vars, lv, err := bo.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, false, err
	}
	vars, rv, err := bo.Right.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, false, err
	}
	return vars, lv, rv, nil
}

// Attributes returns a slice of all attributes used in the evaluation of this node, including a children nodes.
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

// Variables returns a slice of all variables used in the evaluation of this node, including a children nodes.
// Attributes de-duplicates values from the left node and right node using a set.
func (bo BinaryOperator) Variables() []string {
	set := make(map[string]struct{})
	for _, x := range bo.Left.Variables() {
		set[x] = struct{}{}
	}
	for _, x := range bo.Right.Variables() {
		set[x] = struct{}{}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
