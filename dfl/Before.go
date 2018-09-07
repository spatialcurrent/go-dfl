// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Before is a TemporalBinaryOperator evaluating to true if the left value is before the right value.
// The left and right values must be string, time.Time, or *time.Time.
type Before struct {
	*TemporalBinaryOperator // Extends the TemporalBinaryOperator struct
}

func (b Before) Dfl(quotes []string, pretty bool, tabs int) string {
	return b.BinaryOperator.Dfl("before", quotes, pretty, tabs)
}

func (b Before) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "before",
		"left":  b.Left.Map(),
		"right": b.Right.Map(),
	}
}

func (b Before) Compile() Node {
	left := b.Left.Compile()
	right := b.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := CompareTimes(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: (v < 0)}
		}
	}
	return Before{&TemporalBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (b Before) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, v, err := b.EvaluateAndCompare(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}

	return vars, v < 0, nil
}
