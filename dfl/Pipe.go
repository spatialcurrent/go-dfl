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
	"strings"
)

// Pipe is a BinaryOperator which represents the "|" pipe operation of left and right values.
type Pipe struct {
	*BinaryOperator
}

func (p Pipe) Last() Node {
	switch right := p.Right.(type) {
	case Pipe:
		return right.Last()
	}
	return p
}

func (p Pipe) Dfl(quotes []string, pretty bool, tabs int) string {
	if pretty {
		switch p.Left.(type) {
		case *Literal:
			switch p.Left.(*Literal).Value.(type) {
			case string, int, []byte, Null:
				return strings.Repeat("  ", tabs) + p.Left.Dfl(quotes, false, tabs) + " | " + p.Right.Dfl(quotes, false, tabs)
			}
		}
		switch p.Right.(type) {
		case *Literal:
			switch p.Right.(*Literal).Value.(type) {
			case string, int, []byte, Null:
				return strings.Repeat("  ", tabs) + p.Left.Dfl(quotes, false, tabs) + " | " + p.Right.Dfl(quotes, false, tabs)
			}
		}
		return strings.Repeat("  ", tabs) + p.Left.Dfl(quotes, pretty, tabs) + " | " + "\n" + p.Right.Dfl(quotes, pretty, tabs)
	}
	return p.Left.Dfl(quotes, pretty, tabs) + " | " + p.Right.Dfl(quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (p Pipe) Sql(pretty bool, tabs int) string {
	stmt := "SELECT * FROM (SELECT * FROM $TABLE " + p.Left.Sql(pretty, tabs) + ") as A" + fmt.Sprint(tabs) + " WHERE " + p.Right.Sql(pretty, tabs) + ";"
	return stmt
}

func (p Pipe) Map() map[string]interface{} {
	return p.BinaryOperator.Map("pipe", p.Left, p.Right)
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

func (p Pipe) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, lv, err := p.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, lv, errors.Wrap(err, "error processing left value of "+p.Dfl(quotes, false, 0))
	}
	vars, rv, err := p.Right.Evaluate(vars, lv, funcs, quotes)
	if err != nil {
		return vars, rv, errors.Wrap(err, "error processing right value of "+p.Dfl(quotes, false, 0))
	}
	return vars, rv, nil
}
