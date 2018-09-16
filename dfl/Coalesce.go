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

// Sql returns the SQL representation of this node as a string
func (c Coalesce) Sql(pretty bool, tabs int) string {
	return "COALESCE(" + c.Left.Sql(pretty, tabs) + ", " + c.Right.Sql(pretty, tabs) + ")"
}

func (c Coalesce) Map() map[string]interface{} {
	return c.BinaryOperator.Map("?:", c.Left, c.Right)
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

func (c Coalesce) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, lv, err := c.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, lv, errors.Wrap(err, "Error evaluating left value of coalesce: "+c.Left.Dfl(quotes, false, 0))
	}

	switch lv.(type) {
	case Null:
		vars, rv, err := c.Right.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, rv, errors.Wrap(err, "Error evaluating right value of Coalesce: "+c.Left.Dfl(quotes, false, 0))
		}
		return vars, rv, nil
	}

	return vars, lv, nil
}
