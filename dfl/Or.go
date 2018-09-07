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

// Or is a BinaryOperator which represents the logical boolean OR operation of left and right values.
type Or struct {
	*BinaryOperator
}

func (o Or) Dfl(quotes []string, pretty bool, tabs int) string {
	return o.BinaryOperator.Dfl("or", quotes, pretty, tabs)
}

func (o Or) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "or",
		"left":  o.Left.Map(),
		"right": o.Right.Map(),
	}
}

// Compile returns a compiled version of this node.
// If the left value and right value are both compiled as Literals, then returns the logical boolean AND operation of the left and right value.
// Otherwise, returns a clone.
func (o Or) Compile() Node {
	left := o.Left.Compile()
	right := o.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			switch left.(Literal).Value.(type) {
			case bool:
				switch right.(Literal).Value.(type) {
				case bool:
					return Literal{Value: (left.(Literal).Value.(bool) || right.(Literal).Value.(bool))}
				}
			}
		}
	}
	return Or{&BinaryOperator{Left: left, Right: right}}
}

func (o Or) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, lv, err := o.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}
	switch lv.(type) {
	case bool:
		if lv.(bool) {
			return vars, true, nil
		}
		vars, rv, err := o.Right.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, false, err
		}
		switch rv.(type) {
		case bool:
			return vars, rv.(bool), nil
		}
	}
	return vars, false, errors.New("Error evaluating expression " + o.Dfl(quotes, false, 0))
}
