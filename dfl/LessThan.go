// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

type LessThan struct {
	*NumericBinaryOperator
}

func (lt LessThan) Dfl(quotes []string, pretty bool) string {
	return "(" + lt.Left.Dfl(quotes, pretty) + " < " + lt.Right.Dfl(quotes, pretty) + ")"
}

func (lt LessThan) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "<",
		"left":  lt.Left.Map(),
		"right": lt.Right.Map(),
	}
}

func (lt LessThan) Compile() Node {
	left := lt.Left.Compile()
	right := lt.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := CompareNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: (v < 0)}
		}
	}
	return LessThan{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (lt LessThan) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	v, err := lt.EvaluateAndCompare(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}

	return v < 0, nil
}
