// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Divide is a NumericBinaryOperator that represents the mathematical division of two nodes.
type Divide struct {
	*NumericBinaryOperator
}

// Dfl returns the DFL representation of this node as a string
func (d Divide) Dfl(quotes []string, pretty bool, tabs int) string {
	return d.BinaryOperator.Dfl("/", quotes, pretty, tabs)
}

// Map returns a map representation of this node.
func (d Divide) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "/",
		"left":  d.Left.Map(),
		"right": d.Right.Map(),
	}
}

// Compile returns a compiled version of this node.
func (d Divide) Compile() Node {
	left := d.Left.Compile()
	right := d.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := DivideNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: v}
		}
	}
	return &Divide{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

// Evaluate evaluates this node given the variables "vars" and context "ctx" and returns the output, and an error if any.
func (d Divide) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, rv, err := d.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	v, err := DivideNumbers(lv, rv)
	if err != nil {
		return vars, 0, err
	}

	return vars, v, err
}
