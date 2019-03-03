// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"reflect"
)

import (
	"github.com/pkg/errors"
)

// Within is a BinaryOperator that represents that the left value is between
type Within struct {
	*BinaryOperator
}

// Dfl returns the DFL representation of this node as a string
func (w Within) Dfl(quotes []string, pretty bool, tabs int) string {
	return w.BinaryOperator.Dfl("within", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (w Within) Sql(pretty bool, tabs int) string {
	str := w.Left.Sql(pretty, tabs) + " BETWEEN "
	switch right := w.Right.(type) {
	case Literal:
		if t := reflect.TypeOf(right.Value); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
			if rv := reflect.ValueOf(right.Value); rv.Len() == 2 {
				str += fmt.Sprint(rv.Index(0).Interface()) + " AND " + fmt.Sprint(rv.Index(1).Interface())
			}
		}
	case *Attribute:
		str += right.Sql(pretty, tabs) + "[0] AND " + right.Sql(pretty, tabs) + "[1]"
	case *Variable:
		str += right.Sql(pretty, tabs) + "[0] AND " + right.Sql(pretty, tabs) + "[1]"
	case Array:
		if right.Len() == 2 {
			str += right.Nodes[0].Sql(pretty, tabs) + " AND " + right.Nodes[1].Sql(pretty, tabs)
		}
	case And:
		str += right.Left.Sql(pretty, tabs) + " AND " + right.Right.Sql(pretty, tabs)
	case *And:
		str += right.Left.Sql(pretty, tabs) + " AND " + right.Right.Sql(pretty, tabs)
	}
	return str
}

// Map returns a map representation of this node
func (w Within) Map() map[string]interface{} {
	return w.BinaryOperator.Map("within", w.Left, w.Right)
}

// Compile returns a compiled version of this node.
// If the left and right values are both compiled as literals, then returns the compiled Literal with that value set.
// Otherwise returns a clone of this node.
func (w Within) Compile() Node {
	left := w.Left.Compile()
	right := w.Right.Compile()

	switch right.(type) {
	case And:
		right = Array{Nodes: []Node{right.(And).Left, right.(And).Right}}
	}

	switch left.(type) {
	case Literal:
		switch left.(Literal).Value.(type) {
		case int, int8, int16, int32, int64, float64:
			switch right.(type) {
			case Literal:
				if t := reflect.TypeOf(right.(Literal).Value); t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
					rv := reflect.ValueOf(right.(Literal).Value)
					if rv.Len() == 2 {
						v, err := WithinRange(left.(Literal).Value, rv.Index(0).Interface(), rv.Index(1).Interface())
						if err != nil {
							return &Within{&BinaryOperator{Left: left, Right: right}}
						}
						return Literal{Value: v}
					}
				}
			}
		}
	}

	return &Within{&BinaryOperator{Left: left, Right: right}}
}

// Evaluate returns the value of this node given Context ctx, and an error if any.
func (w Within) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, err := w.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating left value for "+w.Dfl(quotes, false, 0))
	}

	vars, rv, err := w.Right.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating right value for "+w.Dfl(quotes, false, 0))
	}

	if t := reflect.TypeOf(rv); !(t.Kind() == reflect.Array || t.Kind() == reflect.Slice) {
		return vars, false, errors.Wrap(err, "right value is wrong type for "+w.Dfl(quotes, false, 0))
	}

	rvv := reflect.ValueOf(rv)
	if rvv.Len() != 2 {
		return vars, false, errors.Wrap(err, "right value is invalid length "+w.Dfl(quotes, false, 0))
	}

	v, err := WithinRange(lv, rvv.Index(0).Interface(), rvv.Index(1).Interface())
	return vars, v, err
}
