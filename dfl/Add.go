// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Add is a NumericBinaryOperator that represents the mathematical addition of two nodes.
type Add struct {
	*NumericBinaryOperator
}

// Dfl returns the DFL representation of this node as a string
func (a Add) Dfl(quotes []string, pretty bool) string {
	return "(" + a.Left.Dfl(quotes, pretty) + " + " + a.Right.Dfl(quotes, pretty) + ")"
}

// Map returns a map representation of this node
func (a Add) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "+",
		"left":  a.Left.Map(),
		"right": a.Right.Map(),
	}
}

// Compile returns a compiled version of this node.
// If the left and right values are both compiled as literals, then returns the compiled Literal with that value set.
// Otherwise returns a clone of this node.
func (a Add) Compile() Node {
	left := a.Left.Compile()
	right := a.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := AddValues(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: v}
		}
	}
	return Add{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

// Evaluate returns the value of this node given Context ctx, and an error if any.
func (a Add) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	lv, rv, err := a.EvaluateLeftAndRight(ctx, funcs, quotes)
	if err != nil {
		return 0, err
	}

	v, err := AddValues(lv, rv)
	if err != nil {
		return 0, err
	}

	return v, err
}
