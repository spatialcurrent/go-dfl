// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// GreaterThan is a NumericBinaryOperator that evaluating to true if parameter a is greater than parameter b.
// The parameters may be of type int, int64, or float64.
type GreaterThan struct {
	*NumericBinaryOperator
}

func (gt GreaterThan) Dfl(quotes []string, pretty bool) string {
	return "(" + gt.Left.Dfl(quotes, pretty) + " > " + gt.Right.Dfl(quotes, pretty) + ")"
}

func (gt GreaterThan) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    ">",
		"left":  gt.Left.Map(),
		"right": gt.Right.Map(),
	}
}

func (gt GreaterThan) Compile() Node {
	left := gt.Left.Compile()
	right := gt.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := CompareNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: (v > 0)}
		}
	}
	return GreaterThan{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (gt GreaterThan) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	v, err := gt.EvaluateAndCompare(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}

	return v > 0, nil
}
