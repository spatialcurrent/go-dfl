// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// GreaterThanOrEqual is a NumericBinaryOperator that evaluating to true if parameter a is greater than or equal to parameter b.
// The parameters may be of type int, int64, or float64.
type GreaterThanOrEqual struct {
	*NumericBinaryOperator
}

func (gte GreaterThanOrEqual) Dfl(quotes []string, pretty bool, tabs int) string {
	return gte.BinaryOperator.Dfl(">=", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (gte GreaterThanOrEqual) Sql(pretty bool, tabs int) string {
	return gte.BinaryOperator.Sql(">=", pretty, tabs)
}

func (gte GreaterThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
		"@type": ">=",
		"@value": map[string]interface{}{
			"left":  gte.Left.Map(),
			"right": gte.Right.Map(),
		},
	}
}

func (gte GreaterThanOrEqual) MarshalMap() (interface{}, error) {
	return gte.Map(), nil
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
			return &Literal{Value: (v >= 0)}
		}
	}
	return &GreaterThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (gte GreaterThanOrEqual) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, v, err := gte.EvaluateAndCompare(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}

	return vars, v >= 0, nil
}
