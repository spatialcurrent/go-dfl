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

func (b Before) Dfl() string {
	return "(" + b.Left.Dfl() + " before " + b.Right.Dfl() + ")"
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

func (b Before) Evaluate(ctx interface{}, funcs FunctionMap) (interface{}, error) {

	v, err := b.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v < 0, nil
}
