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

// Function is a refrenced function in a DFL filter.  The actual function in a given FunctionMap is derefernced by name.
type Function struct {
	Name      string `json:"name" bson:"name" yaml:"name" hcl:"name"`                     // name of the function
	Arguments []Node `json:"arguments" bson:"arguments" yaml:"arguments" hcl:"arguments"` // list of function arguments
}

func (f Function) Dfl(quotes []string, pretty bool) string {
	out := f.Name + "("
	for i, arg := range f.Arguments {
		if i > 0 {
			out += ", "
		}
		out += arg.Dfl(quotes[1:], pretty)
	}
	out += ")"
	return out
}

func (f Function) Compile() Node {
	return f
}

func (f Function) Map() map[string]interface{} {
	arguments := make([]map[string]interface{}, 0, len(f.Arguments))
	for _, a := range f.Arguments {
		arguments = append(arguments, a.Map())
	}
	return map[string]interface{}{
		"op":        "function",
		"name":      f.Name,
		"arguments": arguments,
	}
}

func (f Function) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {
	if fn, ok := funcs[f.Name]; ok {
		values := make([]interface{}, 0, len(f.Arguments))
		for _, arg := range f.Arguments {
			value, err := arg.Evaluate(ctx, funcs, quotes)
			if err != nil {
				return &Null{}, err
			}
			values = append(values, value)
		}
		return fn(funcs, ctx, values, quotes)
	} else {
		return "", errors.New("attempted to evaluate unknown function " + f.Name)
	}
}

func (f Function) Attributes() []string {
	set := make(map[string]struct{})
	for _, n := range f.Arguments {
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
