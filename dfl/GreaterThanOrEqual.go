// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// GreaterThanOrEqual is a NumericBinaryOperator that evaluating to true if parameter a is greater than or equal to parameter b.
// The parameters may be of type int, int64, or float64.
type GreaterThanOrEqual struct {
	*NumericBinaryOperator
}

func (gte GreaterThanOrEqual) Dfl(quotes []string, pretty bool) string {
	return "(" + gte.Left.Dfl(quotes, pretty) + " >= " + gte.Right.Dfl(quotes, pretty) + ")"
}

func (gte GreaterThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    ">=",
		"left":  gte.Left.Map(),
		"right": gte.Right.Map(),
	}
}

func (gte GreaterThanOrEqual) Compile() Node {
	left := gte.Left.Compile()
	right := gte.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := CompareNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: (v >= 0)}
		}
	}
	return GreaterThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (gte GreaterThanOrEqual) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	v, err := gte.EvaluateAndCompare(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}

	return v >= 0, nil
}
