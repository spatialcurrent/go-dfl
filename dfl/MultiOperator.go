// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// MultiOperator represents an operator with a dynamic list of arguments.
type MultiOperator struct {
	Arguments []Node `json:"arguments" bson:"arguments" yaml:"arguments" hcl:"arguments"` // list of function arguments
}

// First returns the last argument for this operator, if exists.
func (mo MultiOperator) First() Node {
	if len(mo.Arguments) > 0 {
		return mo.Arguments[0]
	}
	return nil
}

// Last returns the last argument for this operator, if exists.
func (mo MultiOperator) Last() Node {
	if len(mo.Arguments) > 0 {
		return mo.Arguments[len(mo.Arguments)-1]
	}
	return nil
}

func (mo MultiOperator) Map(operator string) map[string]interface{} {
	arguments := make([]map[string]interface{}, 0, len(mo.Arguments))
	for _, a := range mo.Arguments {
		arguments = append(arguments, a.Map())
	}
	return map[string]interface{}{
		"op":        operator,
		"arguments": arguments,
	}
}

func (mo MultiOperator) Attributes() []string {
	set := make(map[string]struct{})
	for _, n := range mo.Arguments {
		for _, x := range n.Attributes() {
			set[x] = struct{}{}
		}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}

func (mo MultiOperator) Variables() []string {
	set := make(map[string]struct{})
	for _, n := range mo.Arguments {
		for _, x := range n.Variables() {
			set[x] = struct{}{}
		}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
