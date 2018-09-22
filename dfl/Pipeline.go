// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Pipeline is a Node representing a pipeline of nodes where the output of each node is used as the input for the next.
type Pipeline struct {
	Nodes []Node
}

// Len returns the length of the underlying array.
func (p Pipeline) Len() int {
	return len(p.Nodes)
}

// Dfl returns the DFL representation of this node as a string
func (p Pipeline) Dfl(quotes []string, pretty bool, tabs int) string {
	str := ""
	for i, x := range p.Nodes {
		if i > 0 {
			str += " | "
		}
		str += x.Dfl(quotes, pretty, tabs)
	}
	return str
}

// Sql returns the SQL representation of this node as a string
func (p Pipeline) Sql(pretty bool, tabs int) string {
	return ""
}

func (p Pipeline) Map() map[string]interface{} {
	return map[string]interface{}{
		"nodes": p.Nodes,
	}
}

// Compile returns a compiled version of this node.
// If all the values of an Set are literals, returns a single Literal with the corresponding array as its value.
// Otherwise returns the original node..
func (p Pipeline) Compile() Node {
	return Pipeline{Nodes: p.Nodes}
}

func (p Pipeline) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	for _, n := range p.Nodes {
		_, v, err := n.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, v, err
		}
		ctx = v
	}
	return vars, ctx, nil
}

func (p Pipeline) Attributes() []string {
	set := make(map[string]struct{})
	for _, n := range p.Nodes {
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

func (p Pipeline) Variables() []string {
	set := make(map[string]struct{})
	for _, n := range p.Nodes {
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
