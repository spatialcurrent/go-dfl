// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

type Subtract struct {
	*NumericBinaryOperator
}

func (s Subtract) Dfl(quotes []string, pretty bool, tabs int) string {
	return s.BinaryOperator.Dfl("-", quotes, pretty, tabs)
}

func (s Subtract) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "-",
		"left":  s.Left.Map(),
		"right": s.Right.Map(),
	}
}

func (s Subtract) Compile() Node {
	left := s.Left.Compile()
	right := s.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := SubtractNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: v}
		}
	}
	return Subtract{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (s Subtract) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	lv, rv, err := s.EvaluateLeftAndRight(ctx, funcs, quotes)
	if err != nil {
		return 0, err
	}

	v, err := SubtractNumbers(lv, rv)
	if err != nil {
		return 0, err
	}

	return v, err
}
