// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// TemporalBinaryOperator is an abstract struct
// NumericBinaryOperator is a convience struct that attaches to EvaluateAndCompare function
// that is used by structs implementing the Node interface.
type TemporalBinaryOperator struct {
	*BinaryOperator // Extends the BinaryOperator struct
}

// EvaluateAndCompare returns the value of the node given the Context ctx, and error if any.
// If the left value and right value are at the same time, returns 0.
// If the left value is before the right value, returns -1.
// if the left value is after the right value, returns 1.
func (tbo TemporalBinaryOperator) EvaluateAndCompare(ctx Context, funcs FunctionMap) (int, error) {

	lv, rv, err := tbo.EvaluateLeftAndRight(ctx, funcs)
	if err != nil {
		return 0, err
	}

	v, err := CompareTimes(lv, rv)
	if err != nil {
		return 0, err
	}

	return v, err

}
