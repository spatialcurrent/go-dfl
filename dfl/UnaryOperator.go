// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// UnaryOperator is an abstract Node the represents an operator with only 1 operand.
// THe only implementing struct is the "Not" struct.
type UnaryOperator struct {
	Node Node
}

// Attributes returns the context attributes used by the child node, if any.
func (uo UnaryOperator) Attributes() []string {
	return uo.Node.Attributes()
}
