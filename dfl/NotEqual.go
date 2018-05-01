// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// NotEqual is a NumericBinaryOperator that evaluates to true if the left value is not equal to the right value.
// The values may be of type int, int64, or float64.
type NotEqual struct {
	*NumericBinaryOperator
}

// Dfl returns the DFL expression representation of the node as a string value.
// For example
//	"( @amenity  !=  shop )"
func (ne NotEqual) Dfl() string {
	return "(" + ne.Left.Dfl() + " != " + ne.Right.Dfl() + ")"
}

func (ne NotEqual) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "equal",
		"left":  ne.Left.Map(),
		"right": ne.Right.Map(),
	}
}

func (ne NotEqual) Compile() Node {
	return ne
}

func (ne NotEqual) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {

	v, err := ne.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v != 0, nil
}
