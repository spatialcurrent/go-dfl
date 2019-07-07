// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"github.com/spatialcurrent/go-dfl/pkg/dfl/builder"
)

// Print is a UnaryOperator that prints a value to stdout
type Print struct {
	*UnaryOperator
}

// Dfl returns the DFL representation of this node (and its children nodes)
func (p Print) Dfl(quotes []string, pretty bool, tabs int) string {
	b := builder.New(quotes, tabs).Pretty(pretty).Op("print").Right(p.Node)
	if pretty {
		b = b.Indent(tabs)
	}
	return b.Dfl()
}

// Sql returns the SQL representation of this node as a string
func (p Print) Sql(pretty bool, tabs int) string {
	panic("Print is not supported.")
}

// Map returns a map representation of this node (and its children nodes)
func (p Print) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":   "print",
		"node": p.Node.Map(),
	}
}

// Compile returns a compiled version of this node.
// If the the child Node is compiled as a boolean Literal, then returns an inverse Literal Node.
// Otherwise returns a clone of this node.
func (p Print) Compile() Node {
	return Print{&UnaryOperator{Node: p.Node.Compile()}}
}

// Evaluate evaluates this node within a context and returns the bool result, and error if any.
func (p Print) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, v, err := p.Node.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, ctx, err
	}
	fmt.Println(TryFormatLiteral(v, quotes, false, 0))
	return vars, ctx, nil
}
