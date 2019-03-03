// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"reflect"
)

// Set is a Node representing a set of values, which can be either a Literal or Attribute.
type Set struct {
	Nodes []Node
}

// Len returns the length of the underlying array.
func (s Set) Len() int {
	return len(s.Nodes)
}

func (s Set) Dfl(quotes []string, pretty bool, tabs int) string {
	str := "{"
	for i, x := range s.Nodes {
		if i > 0 {
			str += ", "
		}
		str += x.Dfl(quotes, pretty, tabs)
	}
	str = str + "}"
	return str
}

// Sql returns the SQL representation of this node as a string
func (s Set) Sql(pretty bool, tabs int) string {
	str := SqlArrayPrefix
	for i, x := range s.Nodes {
		if i > 0 {
			str += ", "
		}
		str += x.Sql(pretty, tabs)
	}
	str = str + SqlArraySuffix
	return str
}

func (a Set) Map() map[string]interface{} {
	return map[string]interface{}{
		"nodes": a.Nodes,
	}
}

// Compile returns a compiled version of this node.
// If all the values of an Set are literals, returns a single Literal with the corresponding Set/slice as its value.
// Otherwise returns the original node..
func (a Set) Compile() Node {
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
	set := NewStringSet()
	switch arr := TryConvertArray(values).(type) {
	case []string:
		set.Add(arr...)
	case []interface{}:
		for _, x := range arr {
			set.Add(fmt.Sprint(x))
		}
	}
	return Literal{Value: map[string]struct{}(set)}
}

func (a Set) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	values := make([]interface{}, len(a.Nodes))
	for i, n := range a.Nodes {
		vars, v, err := n.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, values, err
		}
		values[i] = v
	}

	set := NewStringSet()
	switch arr := TryConvertArray(values).(type) {
	case []string:
		set.Add(arr...)
	case []interface{}:
		for _, x := range arr {
			set.Add(fmt.Sprint(x))
		}
	}

	return vars, map[string]struct{}(set), nil
}

func (a Set) Attributes() []string {
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

func (a Set) Variables() []string {
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
