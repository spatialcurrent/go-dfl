// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// TestCase is a struct containing the variables for a unit test of expression evaluation.
type TestCase struct {
	Expression string      // the DFL expression
	Context    interface{} // the Context to use for evaluation
	Result     interface{} // The result of the evaluation
}

// NewTestCase returns a new TestCase
func NewTestCase(exp string, ctx interface{}, result interface{}) TestCase {
	return TestCase{
		Expression: exp,
		Context:    ctx,
		Result:     result,
	}
}
