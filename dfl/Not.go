// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
)

// Not is a UnaryOperator that inverts the boolean value of the children Node.
type Not struct {
	*UnaryOperator
}

// Dfl returns the DFL representation of this node (and its children nodes)
func (n Not) Dfl(quotes []string, pretty bool) string {
	return "not " + n.Node.Dfl(quotes, pretty)
}

// Map returns a map representation of this node (and its children nodes)
func (n Not) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":   "not",
		"node": n.Node.Map(),
	}
}

// Compile returns a compiled version of this node.
// If the the child Node is compiled as a boolean Literal, then returns an inverse Literal Node.
// Otherwise returns a clone of this node.
func (n Not) Compile() Node {
	child := n.Node.Compile()
	switch child.(type) {
	case Literal:
		switch child.(Literal).Value.(type) {
		case bool:
			return Literal{Value: !child.(Literal).Value.(bool)}
		}
	}
	return Not{&UnaryOperator{Node: child}}
}

// Evaluate evaluates this node within a context and returns the bool result, and error if any.
func (n Not) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {
	v, err := n.Node.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}
	switch v.(type) {
	case bool:
		return !(v.(bool)), nil
	}
	return false, errors.New("Error evaluating expression " + n.Dfl(quotes, false))
}
