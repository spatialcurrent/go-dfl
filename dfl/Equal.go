// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Equal is a NumericBinaryOperator that evaluating to true if parameter a is equal to parameter b.
// The parameters may be of type int, int64, or float64.
type Equal struct {
	*NumericBinaryOperator
}

func (e Equal) Dfl() string {
	return "(" + e.Left.Dfl() + " == " + e.Right.Dfl() + ")"
}

func (e Equal) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "equal",
		"left":  e.Left.Map(),
		"right": e.Right.Map(),
	}
}

func (e Equal) Compile() Node {
	return e
}

func (e Equal) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {

	v, err := e.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v == 0, nil
}
