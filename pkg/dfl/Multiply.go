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

// Multiply is a NumericBinaryOperator that represents the mathematical multiplication of two nodes.
type Multiply struct {
	*NumericBinaryOperator
}

// Dfl returns the DFL representation of this node as a string
func (m Multiply) Dfl(quotes []string, pretty bool, tabs int) string {
	return m.BinaryOperator.Dfl("*", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (m Multiply) Sql(pretty bool, tabs int) string {
	return m.BinaryOperator.Sql("*", pretty, tabs)
}

// Map returns a map representation of this node.
func (m Multiply) Map() map[string]interface{} {
	return m.BinaryOperator.Map("multiply", m.Left, m.Right)
}

// Compile returns a compiled version of this node.
func (m Multiply) Compile() Node {
	left := m.Left.Compile()
	right := m.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := af.Multiply.ValidateRun([]interface{}{left.(Literal).Value, right.(Literal).Value})
			if err != nil {
				return &Multiply{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
			}
			return Literal{Value: v}
		}
	}
	return &Multiply{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

// Evaluate evaluates this node given the variables "vars" and context "ctx" and returns the output, and an error if any.
func (m Multiply) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, rv, err := m.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	v, err := af.Multiply.ValidateRun(lv, rv)
	if err != nil {
		return vars, 0, errors.Wrap(err, ErrorEvaluate{Node: m, Quotes: quotes}.Error())
	}

	return vars, v, err
}
