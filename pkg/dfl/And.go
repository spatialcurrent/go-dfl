// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
)

// And is a BinaryOperator which represents the logical boolean AND operation of left and right values.
type And struct {
	*BinaryOperator
}

func (a And) Dfl(quotes []string, pretty bool, tabs int) string {
	return a.BinaryOperator.Dfl("and", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (a And) Sql(pretty bool, tabs int) string {
	return a.BinaryOperator.Sql("and", pretty, tabs)
}

func (a And) Map() map[string]interface{} {
	return a.BinaryOperator.Map("and", a.Left, a.Right)
}

// Compile returns a compiled version of this node.
// If the left value and right value are both compiled as Literals, then returns the logical boolean AND operation of the left and right value.
// Otherwise, returns a clone.
func (a And) Compile() Node {
	left := a.Left.Compile()
	right := a.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			switch left.(Literal).Value.(type) {
			case bool:
				switch right.(Literal).Value.(type) {
				case bool:
					return Literal{Value: (left.(Literal).Value.(bool) && right.(Literal).Value.(bool))}
				}
			}
		}
	}
	return And{&BinaryOperator{Left: left, Right: right}}
}

func (a And) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, lv, err := a.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "error evaluating left node")
	}
	switch lv.(type) {
	case bool:
		if !lv.(bool) {
			return vars, false, nil
		}
		vars, rv, err := a.Right.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, false, errors.Wrap(err, "error evaluating right node")
		}
		switch rv.(type) {
		case bool:
			return vars, rv.(bool), nil
		}
		return vars, false, errors.Wrap(&ErrorEvaluate{Node: a, Quotes: quotes}, "right value is not a bool")
	}
	return vars, false, errors.Wrap(&ErrorEvaluate{Node: a, Quotes: quotes}, "left value is not a bool")
}
