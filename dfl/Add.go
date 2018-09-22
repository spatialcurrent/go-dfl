// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-adaptive-functions/af"
)

// Add is a BinaryOperator that represents the addition of two nodes.
type Add struct {
	*BinaryOperator
}

// Dfl returns the DFL representation of this node as a string
func (a Add) Dfl(quotes []string, pretty bool, tabs int) string {
	return a.BinaryOperator.Dfl("+", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node.
func (a Add) Sql(pretty bool, tabs int) string {
	return a.BinaryOperator.Sql("+", pretty, tabs)
}

// Map returns a map representation of this node.
func (a Add) Map() map[string]interface{} {
	return a.BinaryOperator.Map("+", a.Left, a.Right)
}

// Compile returns a compiled version of this node.
// If the left and right values are both compiled as literals, then returns the compiled Literal with that value set.
// Otherwise returns a clone of this node.
func (a Add) Compile() Node {
	left := a.Left.Compile()
	right := a.Right.Compile()

	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := af.Add.ValidateRun([]interface{}{left.(Literal).Value, right.(Literal).Value})
			if err != nil {
				return &Add{&BinaryOperator{Left: left, Right: right}}
			}
			return Literal{Value: v}
		}
		switch left.(Literal).Value.(type) {
		case string:
			switch right.(type) {
			case Concat:
				switch right.(Concat).Arguments[0].(type) {
				case Literal:
					n := Literal{
						Value: left.(Literal).Value.(string) + fmt.Sprint(right.(Concat).Arguments[0].(Literal).Value),
					}
					return Concat{&MultiOperator{Arguments: append([]Node{n}, right.(Concat).Arguments[1:]...)}}
				}
				return Concat{&MultiOperator{Arguments: append([]Node{left}, right.(Concat).Arguments...)}}
			}
			return Concat{&MultiOperator{Arguments: []Node{left, right}}}
		}
	case Attribute, *Attribute, Variable, *Variable:
		switch right.(type) {
		case Literal:
			switch right.(Literal).Value.(type) {
			case string:
				return Concat{&MultiOperator{Arguments: []Node{left, right}}}
			}
		case Concat:
			return Concat{&MultiOperator{Arguments: append([]Node{left}, right.(Concat).Arguments...)}}
		}
	}
	return &Add{&BinaryOperator{Left: left, Right: right}}
}

// Evaluate returns the value of this node given Context ctx, and an error if any.
func (a Add) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, rv, err := a.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	v, err := af.Add.ValidateRun([]interface{}{lv, rv})
	if err != nil {
		return vars, 0, errors.Wrap(err, ErrorEvaluate{Node: a, Quotes: quotes}.Error())
	}

	return vars, v, err
}
