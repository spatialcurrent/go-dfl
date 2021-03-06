// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

// Dictionary is a Node representing a dictionary of key value pairs.
type Dictionary struct {
	Nodes map[Node]Node
}

func NewDictionary(m map[string]interface{}) *Dictionary {
	nodes := map[Node]Node{}
	for k, v := range m {
		if d, ok := v.(map[string]interface{}); ok {
			nodes[&Literal{Value: k}] = NewDictionary(d)
		} else {
			nodes[&Literal{Value: k}] = &Literal{Value: v}
		}
	}
	return &Dictionary{Nodes: nodes}
}

// Len returns the length of the underlying array.
func (d Dictionary) Len() int {
	return len(d.Nodes)
}

func (d Dictionary) Dfl(quotes []string, pretty bool, tabs int) string {
	if len(d.Nodes) == 0 {
		return "{}"
	}
	values := make([]string, 0)
	for k, v := range d.Nodes {
		values = append(values, k.Dfl(quotes, pretty, tabs+1)+": "+v.Dfl(quotes, pretty, tabs+1))
	}
	if pretty {
		return "{" + "\n" + FormatList(values, ",", pretty, tabs+1) + "\n" + strings.Repeat(DefaultTab, tabs) + "}"
	}
	return "{" + FormatList(values, ",", pretty, tabs) + "}"

}

// Sql returns the SQL representation of this node as a string
func (d Dictionary) Sql(pretty bool, tabs int) string {
	str := SqlQuote + SqlArrayPrefix
	i := 0
	for k, v := range d.Nodes {
		if i > 0 {
			str += ", "
		}
		str += k.Sql(pretty, tabs) + ":" + v.Sql(pretty, tabs)
		i += 1
	}
	str = str + SqlArraySuffix + SqlQuote + "::json"
	return str
}

func (d Dictionary) Map() map[string]interface{} {
	return map[string]interface{}{
		"nodes": d.Nodes,
	}
}

// Compile returns a compiled version of this node.
func (d Dictionary) Compile() Node {
	nodes := map[Node]Node{}
	for k, v := range d.Nodes {
		nodes[k.Compile()] = v.Compile()
	}
	return Dictionary{Nodes: nodes}
}

func (d Dictionary) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	values := map[interface{}]interface{}{}
	for k, v := range d.Nodes {
		_, keyValue, err := k.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, values, err
		}
		_, valueValue, err := v.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, values, err
		}
		values[keyValue] = valueValue
	}
	return vars, values, nil
}

func (d Dictionary) Attributes() []string {
	set := make(map[string]struct{})
	for k, v := range d.Nodes {
		for _, x := range k.Attributes() {
			set[x] = struct{}{}
		}
		for _, x := range v.Attributes() {
			set[x] = struct{}{}
		}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}

func (d Dictionary) Variables() []string {
	set := make(map[string]struct{})
	for k, v := range d.Nodes {
		for _, x := range k.Variables() {
			set[x] = struct{}{}
		}
		for _, x := range v.Variables() {
			set[x] = struct{}{}
		}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
