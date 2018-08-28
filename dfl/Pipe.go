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

// Pipe is a BinaryOperator which represents the "|" pipe operation of left and right values.
type Pipe struct {
	*BinaryOperator
}

func (p Pipe) Dfl(quotes []string, pretty bool, tabs int) string {
	return p.BinaryOperator.Dfl("|", quotes, pretty, tabs)
}

func (p Pipe) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "pipe",
		"left":  p.Left.Map(),
		"right": p.Right.Map(),
	}
}

// Compile returns a compiled version of this node.
// If the left value and right value are both compiled as Literals, then returns the logical boolean AND operation of the left and right value.
// Otherwise, returns a clone.
func (p Pipe) Compile() Node {
	left := p.Left.Compile()
	right := p.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			return Literal{Value: right.(Literal).Value}
		}
	}
	return Pipe{&BinaryOperator{Left: left, Right: right}}
}

func (p Pipe) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {
	lv, err := p.Left.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return lv, errors.Wrap(err, "error processing left value of "+p.Dfl(quotes, false, 0))
	}
	rv, err := p.Right.Evaluate(lv, funcs, quotes)
	if err != nil {
		return rv, errors.Wrap(err, "error processing right value of "+p.Dfl(quotes, false, 0))
	}
	return rv, nil
}
