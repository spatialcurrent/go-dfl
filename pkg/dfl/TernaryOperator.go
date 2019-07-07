// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
	"strings"
)

// TernaryOperator is a DFL Node that represents the ternary operator of a condition, true value, and false value.
type TernaryOperator struct {
	Left  Node
	True  Node
	False Node
}

func (to TernaryOperator) Dfl(quotes []string, pretty bool, tabs int) string {
	if pretty {
		return "(\n" +
			to.Left.Dfl(quotes, pretty, tabs+1) + " ? " +
			to.True.Dfl(quotes, pretty, tabs+1) + " : " +
			to.False.Dfl(quotes, pretty, tabs+2) + "\n" +
			strings.Repeat(DefaultTab, tabs) + ")"
	}
	return "(" + to.Left.Dfl(quotes, pretty, tabs) + " ? " + to.True.Dfl(quotes, pretty, tabs) + " : " + to.False.Dfl(quotes, pretty, tabs) + ")"
}

func (to TernaryOperator) Sql(pretty bool, tabs int) string {
	return "( CASE " + to.Left.Sql(pretty, tabs) + " WHEN true THEN " + to.True.Sql(pretty, tabs) + " ELSE " + to.False.Sql(pretty, tabs) + " END )"
}

func (to TernaryOperator) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":        "ternary",
		"condition": to.Left.Map(),
		"true":      to.True.Map(),
		"false":     to.False.Map(),
	}
}

func (to TernaryOperator) Compile() Node {
	left := to.Left.Compile()
	switch left.(type) {
	case Literal:
		switch left.(Literal).Value.(type) {
		case bool:
			if left.(Literal).Value.(bool) {
				return to.True.Compile()
			} else {
				return to.False.Compile()
			}
		}
	}

	return &TernaryOperator{
		Left:  left,
		True:  to.True.Compile(),
		False: to.False.Compile(),
	}
}

func (to TernaryOperator) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, cv, err := to.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, cv, errors.Wrap(err, "Error evaluating condition of ternary operator: "+to.Left.Dfl(quotes, false, 0))
	}

	switch cv.(type) {
	case bool:
	default:
		return vars, cv, errors.Wrap(err, "ternary operator condition returned a non boolean: "+to.Left.Dfl(quotes, false, 0))
	}

	if cv.(bool) {

		vars, v, err := to.True.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, v, errors.Wrap(err, "Error evaluating true expression of ternary operator: "+to.True.Dfl(quotes, false, 0))
		}

		return vars, v, err
	}

	vars, v, err := to.False.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, v, errors.Wrap(err, "Error evaluating false expression of ternary operator: "+to.False.Dfl(quotes, false, 0))
	}

	return vars, v, err
}

// Attributes returns a slice of all attributes used in the evaluation of this node, including a children nodes.
// Attributes de-duplicates values from the condition, true, and false nodes using a set.
func (to TernaryOperator) Attributes() []string {
	set := make(map[string]struct{})
	for _, x := range to.Left.Attributes() {
		set[x] = struct{}{}
	}
	for _, x := range to.True.Attributes() {
		set[x] = struct{}{}
	}
	for _, x := range to.False.Attributes() {
		set[x] = struct{}{}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}

// Variables returns a slice of all variables used in the evaluation of this node, including a children nodes.
// Variables de-duplicates values from the condition, true, and false nodes using a set.
func (to TernaryOperator) Variables() []string {
	set := make(map[string]struct{})
	for _, x := range to.Left.Variables() {
		set[x] = struct{}{}
	}
	for _, x := range to.True.Variables() {
		set[x] = struct{}{}
	}
	for _, x := range to.False.Variables() {
		set[x] = struct{}{}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
