// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// LessThanOrEqual is a NumericBinaryOperator that evaluating to true if parameter a is less than or equal to parameter b.
// The parameters may be of type int, int64, or float64.
type LessThanOrEqual struct {
	*NumericBinaryOperator
}

func (lte LessThanOrEqual) Dfl(quotes []string, pretty bool, tabs int) string {
	return lte.BinaryOperator.Dfl("<=", quotes, pretty, tabs)
}

func (lte LessThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "<=",
		"left":  lte.Left.Map(),
		"right": lte.Right.Map(),
	}
}

func (lte LessThanOrEqual) Compile() Node {
	left := lte.Left.Compile()
	right := lte.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := CompareNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: (v <= 0)}
		}
	}
	return LessThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (lte LessThanOrEqual) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	v, err := lte.EvaluateAndCompare(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}

	return v <= 0, nil
}
