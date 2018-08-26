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

func (a After) Dfl(quotes []string, pretty bool) string {
	return "(" + a.Left.Dfl(quotes, pretty) + " after " + a.Right.Dfl(quotes, pretty) + ")"
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

func (a After) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	v, err := a.EvaluateAndCompare(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}

	return v > 0, nil
}
