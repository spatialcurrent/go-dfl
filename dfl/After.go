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

func (a After) Dfl(quotes []string, pretty bool, tabs int) string {
	return a.BinaryOperator.Dfl("after", quotes, pretty, tabs)
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

func (a After) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, v, err := a.EvaluateAndCompare(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}

	return vars, v > 0, nil
}
