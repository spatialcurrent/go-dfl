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

// Xor is a BinaryOperator which represents the logical boolean XOR operation of left and right values.
//
//	- https://en.wikipedia.org/wiki/Exclusive_or
type Xor struct {
	*BinaryOperator // Extends the BinaryOperator struct
}

func (x Xor) Dfl() string {
	return "(" + x.Left.Dfl() + " xor " + x.Right.Dfl() + ")"
}

func (x Xor) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "xor",
		"left":  x.Left.Map(),
		"right": x.Right.Map(),
	}
}

// Compile returns a compiled version of this node.
// If the left value and right value are both compiled as Literals, then returns the logical boolean XOR operation of the left and right value.
// Otherwise, returns a clone.
func (x Xor) Compile() Node {
	left := x.Left.Compile()
	right := x.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			switch left.(Literal).Value.(type) {
			case bool:
				switch right.(Literal).Value.(type) {
				case bool:
					return Literal{Value: (left.(Literal).Value.(bool) != right.(Literal).Value.(bool))}
				}
			}
		}
	}
	return Xor{&BinaryOperator{Left: left, Right: right}}
}

func (x Xor) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {
	lv, err := x.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	switch lv.(type) {
	case bool:
		rv, err := x.Right.Evaluate(ctx, funcs)
		if err != nil {
			return false, err
		}
		switch rv.(type) {
		case bool:
			return lv.(bool) != rv.(bool), nil
		}
	}
	return false, errors.New("Error evaulating expression " + x.Dfl())
}
