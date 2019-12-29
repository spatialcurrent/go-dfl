// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
	"strings"
)

// Array is a Node representing an array of values, which can be either a Literal or Attribute.
type Array struct {
	Nodes []Node
}

// Len returns the length of the underlying array.
func (a Array) Len() int {
	return len(a.Nodes)
}

// Dfl returns the DFL representation of this node as a string
func (a Array) Dfl(quotes []string, pretty bool, tabs int) string {
	if len(a.Nodes) == 0 {
		return "[]"
	}
	if len(a.Nodes) == 1 {
		return "[" + strings.TrimSpace(a.Nodes[0].Dfl(quotes, pretty, tabs)) + "]"
	}
	if pretty {
		return "[" + "\n" + FormatList(FormatNodes(a.Nodes, quotes, pretty, tabs), ",", pretty, tabs+1) + "\n" + strings.Repeat(DefaultTab, tabs) + "]"
	}
	return "[" + FormatList(FormatNodes(a.Nodes, quotes, pretty, tabs), ",", pretty, tabs) + "]"
}

// Sql returns the SQL representation of this node as a string
func (a Array) Sql(pretty bool, tabs int) string {
	str := "{"
	for i, x := range a.Nodes {
		if i > 0 {
			str += ", "
		}
		str += x.Sql(pretty, tabs)
	}
	str = str + "}"
	return str
}

func (a Array) Map() map[string]interface{} {
	return map[string]interface{}{
		"@type": "array",
		"@value": map[string]interface{}{
			"nodes": a.Nodes,
		},
	}
}

func (a Array) MarshalMap() (interface{}, error) {
	return a.Map(), nil
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

func (a Array) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	values := make([]interface{}, 0, len(a.Nodes))
	for _, n := range a.Nodes {
		vars, v, err := n.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, values, err
		}
		values = append(values, v)
	}
	return vars, TryConvertArray(values), nil
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

func (a Array) Variables() []string {
	set := make(map[string]struct{})
	for _, n := range a.Nodes {
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
