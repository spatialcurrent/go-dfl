// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
)

// Array is a Node representing an array of values, which can be either a Literal or Attribute.
type Array struct {
	Nodes []Node
}

func (a Array) Dfl() string {
	str := "["
	for i, x := range a.Nodes {
		if i > 0 {
			str += ", "
		}
		str += x.Dfl()
	}
	str = str + "]"
	return str
}

func (a Array) Map() map[string]interface{} {
	return map[string]interface{}{
		"nodes": a.Nodes,
	}
}

// Compile returns a compiled version of this node.
// If all the values of an Set are literals, returns a single Literal with the corresponding array as its value.
// Otherwise returns the original node..
func (a Array) Compile() Node {
	values := make([]interface{}, len(a.Nodes))
	nodes := reflect.ValueOf(a.Nodes)
	for i := 0; i < nodes.Len(); i++ {
		n := nodes.Index(i).Interface()
		switch n.(type) {
		case *Literal:
			values[i] = n.(*Literal).Value
		default:
			return a
		}
	}
	return Literal{Value: TryConvertArray(values)}
}

func (a Array) Evaluate(ctx interface{}, funcs FunctionMap) (interface{}, error) {
	values := make([]interface{}, len(a.Nodes))
	for i, n := range a.Nodes {
		v, err := n.Evaluate(ctx, funcs)
		if err != nil {
			return values, err
		}
		values[i] = v
	}
	return TryConvertArray(values), nil
}

func (a Array) Attributes() []string {
	set := make(map[string]struct{})
	for _, n := range a.Nodes {
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
