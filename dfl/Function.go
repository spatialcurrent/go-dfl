// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
	"strings"
)

// Function is a refrenced function in a DFL filter.  The actual function in a given FunctionMap is derefernced by name.
type Function struct {
	*MultiOperator
	Name string `json:"name" bson:"name" yaml:"name" hcl:"name"` // name of the function
}

func (f Function) Dfl(quotes []string, pretty bool, tabs int) string {
	if pretty {
		if len(f.Arguments) > 0 {
			if len(f.Arguments) == 1 {
				switch arg := f.Arguments[0].(type) {
				case *Attribute:
					return strings.Repeat("  ", tabs) + f.Name + "(" + arg.Dfl(quotes, false, tabs+1) + ")"
				case *Function:
					if len(arg.Arguments) == 0 {
						return strings.Repeat("  ", tabs) + f.Name + "(" + arg.Dfl(quotes, false, tabs+1) + ")"
					} else if len(arg.Arguments) == 1 {
						switch arg.Arguments[0].(type) {
						case *Attribute:
							return strings.Repeat("  ", tabs) + f.Name + "(" + arg.Dfl(quotes, false, tabs+1) + ")"
						}
					}
				}
			}
			out := strings.Repeat("  ", tabs) + f.Name + "("
			for i, arg := range f.Arguments {
				out += "\n" + arg.Dfl(quotes, pretty, tabs+1)
				if i < len(f.Arguments)-1 {
					out += ", "
				} else {
					out += "\n"
				}
			}
			out += strings.Repeat("  ", tabs)
			out += ")"
			return out
		}
		return strings.Repeat("  ", tabs) + f.Name + "()"
	}

	return f.Name + "(" + FormatNodes(f.Arguments, ", ", quotes, pretty, tabs) + ")"
}

// Sql returns the SQL representation of this node as a string
func (f Function) Sql(pretty bool, tabs int) string {
	out := f.Name + "("
	for i, arg := range f.Arguments {
		if i > 0 {
			out += ", "
		}
		out += arg.Sql(pretty, tabs)
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

func (f Function) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	if fn, ok := funcs[f.Name]; ok {
		values := make([]interface{}, 0, len(f.Arguments))
		for _, arg := range f.Arguments {
			_, value, err := arg.Evaluate(vars, ctx, funcs, quotes)
			if err != nil {
				return vars, &Null{}, err
			}
			values = append(values, value)
		}
		v, err := fn(funcs, vars, ctx, values, quotes)
		return vars, v, err
	} else {
		return vars, "", errors.New("attempted to evaluate unknown function " + f.Name)
	}
}
