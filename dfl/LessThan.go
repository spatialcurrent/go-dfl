// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// LessThan is a NumericBinaryOperator that evaluating to true if parameter a is less than parameter b.
// The parameters may be of type int, int64, or float64.
type LessThan struct {
	*NumericBinaryOperator
}

// Sql returns the DFL representation of this node as a string
func (lt LessThan) Dfl(quotes []string, pretty bool, tabs int) string {
	return lt.BinaryOperator.Dfl("<", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (lt LessThan) Sql(pretty bool, tabs int) string {
	return lt.BinaryOperator.Sql("<", pretty, tabs)
}

func (lt LessThan) Map() map[string]interface{} {
	return lt.BinaryOperator.Map("<", lt.Left, lt.Right)
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

func (lt LessThan) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, v, err := lt.EvaluateAndCompare(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}

	return vars, v < 0, nil
}
