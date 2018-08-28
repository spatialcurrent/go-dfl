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

// Coalesce is a BinaryOperator which returns the left value if not null otherwise the right value.
type Coalesce struct {
	*BinaryOperator
}

func (c Coalesce) Dfl(quotes []string, pretty bool, tabs int) string {
	return "(" + c.Left.Dfl(quotes, pretty, tabs) + " ?: " + c.Right.Dfl(quotes, pretty, tabs) + ")"
}

func (c Coalesce) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "?:",
		"left":  c.Left.Map(),
		"right": c.Right.Map(),
	}
}

// Compile returns a compiled version of this node.
// If the left value is compiled as a Literal, then returns the left value.
// Otherwise, returns a clone.
func (c Coalesce) Compile() Node {
	left := c.Left.Compile()
	switch left.(type) {
	case Literal:
		return Literal{Value: left.(Literal).Value}
	}
	right := c.Right.Compile()
	return Coalesce{&BinaryOperator{Left: left, Right: right}}
}

func (c Coalesce) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {
	lv, err := c.Left.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return lv, errors.Wrap(err, "Error evaluating left value of coalesce: "+c.Left.Dfl(quotes, false, 0))
	}

	switch lv.(type) {
	case Null:
		rv, err := c.Right.Evaluate(ctx, funcs, quotes)
		if err != nil {
			return rv, errors.Wrap(err, "Error evaluating right value of Coalesce: "+c.Left.Dfl(quotes, false, 0))
		}
		return rv, nil
	}

	return lv, nil
}
