// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// After is a TemporalBinaryOperator evaluating to true if the left value is after the right value.
// The left and right values must be string, time.Time, or *time.Time.
type After struct {
	*TemporalBinaryOperator // Extends the TemporalBinaryOperator struct
}

func (a After) Dfl() string {
	return "(" + a.Left.Dfl() + " after " + a.Right.Dfl() + ")"
}

func (a After) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "after",
		"left":  a.Left.Map(),
		"right": a.Right.Map(),
	}
}

func (a After) Compile() Node {
	left := a.Left.Compile()
	right := a.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := CompareTimes(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: (v > 0)}
		}
	}
	return After{&TemporalBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (a After) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {

	v, err := a.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v > 0, nil
}
