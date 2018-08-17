// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

type LessThanOrEqual struct {
	*NumericBinaryOperator
}

func (lte LessThanOrEqual) Dfl() string {
	return "(" + lte.Left.Dfl() + " <= " + lte.Right.Dfl() + ")"
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

func (lte LessThanOrEqual) Evaluate(ctx interface{}, funcs FunctionMap) (interface{}, error) {

	v, err := lte.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v <= 0, nil
}
