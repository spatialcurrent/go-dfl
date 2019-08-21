// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
)

// Divide is a NumericBinaryOperator that represents the mathematical division of two nodes.
type Divide struct {
	*NumericBinaryOperator
}

// Dfl returns the DFL representation of this node as a string
func (d Divide) Dfl(quotes []string, pretty bool, tabs int) string {
	return d.BinaryOperator.Dfl("/", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (d Divide) Sql(pretty bool, tabs int) string {
	return d.BinaryOperator.Sql("/", pretty, tabs)
}

// Map returns a map representation of this node.
func (d Divide) Map() map[string]interface{} {
	return d.BinaryOperator.Map("divide", d.Left, d.Right)
}

// Compile returns a compiled version of this node.
func (d Divide) Compile() Node {
	left := d.Left.Compile()
	right := d.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := af.Divide.ValidateRun([]interface{}{left.(Literal).Value, right.(Literal).Value})
			if err != nil {
				return &Divide{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
			}
			return Literal{Value: v}
		}
	}
	return &Divide{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

// Evaluate evaluates this node given the variables "vars" and context "ctx" and returns the output, and an error if any.
func (d Divide) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, rv, err := d.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	v, err := af.Divide.ValidateRun(lv, rv)
	if err != nil {
		return vars, 0, errors.Wrap(err, ErrorEvaluate{Node: d, Quotes: quotes}.Error())
	}

	return vars, v, err
}
