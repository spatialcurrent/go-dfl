// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// NumericBinaryOperator is an abstract struct extending the BinaryOperator struct.
// NumericBinaryOperator is a convience struct that attaches to EvaluateAndCompare function
// that is used by structs implementing the Node interface.
type NumericBinaryOperator struct {
	*BinaryOperator // Extends the BinaryOperator struct
}

// EvaluateAndCompare returns the value of the node given the Context ctx, and error if any.
// If the left value and right value are mathematically equal, returns 0.
// If the left value is less than the right value, returns -1.
// if the left value is greater than the right value, returns 1.
func (nbo NumericBinaryOperator) EvaluateAndCompare(ctx interface{}, funcs FunctionMap, quotes []string) (int, error) {

	lv, rv, err := nbo.EvaluateLeftAndRight(ctx, funcs, quotes)
	if err != nil {
		return 0, err
	}

	v, err := CompareNumbers(lv, rv)
	if err != nil {
		return 0, err
	}

	return v, err

}
