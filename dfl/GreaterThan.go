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

func (gt GreaterThan) Dfl(quotes []string, pretty bool, tabs int) string {
	return gt.BinaryOperator.Dfl(">", quotes, pretty, tabs+1)
}

// Sql returns the SQL representation of this node as a string
func (gt GreaterThan) Sql(pretty bool, tabs int) string {
	return gt.BinaryOperator.Sql(">", pretty, tabs)
}

func (gt GreaterThan) Map() map[string]interface{} {
	return gt.BinaryOperator.Map(">", gt.Left, gt.Right)
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

func (gt GreaterThan) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, v, err := gt.EvaluateAndCompare(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}

	return vars, v > 0, nil
}
