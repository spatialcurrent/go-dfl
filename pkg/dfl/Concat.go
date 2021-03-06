// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
)

// Concat concatenates the string representation of all the arguments
type Concat struct {
	*MultiOperator
}

// Suffix returns the suffix of the result of evaluation, if the last argument is a Literal.  If the last argument is not a literal, then returns an empty string.
func (c Concat) Suffix() string {
	if len(c.Arguments) > 0 {
		switch last := c.Arguments[len(c.Arguments)-1].(type) {
		case Literal:
			return fmt.Sprint(last.Value)
		}
	}
	return ""
}

// Dfl returns the DFL expression representation of this node.
func (c Concat) Dfl(quotes []string, pretty bool, tabs int) string {
	/*
		if pretty {
			if len(c.Arguments) > 0 {
				if len(c.Arguments) == 1 {
					switch arg := c.Arguments[0].(type) {
					case *Attribute:
						return arg.Dfl(quotes, false, tabs)
					case *Function:
						if len(arg.Arguments) == 0 {
							return arg.Dfl(quotes, false, tabs)
						} else if len(arg.Arguments) == 1 {
							switch arg.Arguments[0].(type) {
							case *Attribute:
								return arg.Dfl(quotes, false, tabs)
							}
						}
					}
				}
				out := strings.Repeat("  ", tabs) + "("
				for i, arg := range c.Arguments {
					out += "\n" + arg.Dfl(quotes, pretty, tabs+1)
					if i < len(c.Arguments)-1 {
						out += " + "
					} else {
						out += "\n"
					}
				}
				out += strings.Repeat("  ", tabs)
				out += ")"
				return out
			}
			return strings.Repeat("  ", tabs) + Null{}.Dfl()
		}*/

	return "(" + FormatList(FormatNodes(c.Arguments, quotes, pretty, tabs), "+", pretty, tabs) + ")"
}

// Sql returns the SQL representation of this node as a string
func (c Concat) Sql(pretty bool, tabs int) string {
	out := "concat("
	for i, arg := range c.Arguments {
		if i > 0 {
			out += ", "
		}
		out += arg.Sql(pretty, tabs)
	}
	out += ")"
	return out
}

func (c Concat) Compile() Node {
	if len(c.Arguments) == 0 {
		return &Literal{Value: ""}
	} else if len(c.Arguments) == 1 {
		return c.Arguments[0]
	}
	return Concat{&MultiOperator{Arguments: c.Arguments}}
}

func (c Concat) Map() map[string]interface{} {
	return c.MultiOperator.Map("concat")
}

func (c Concat) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	str := ""
	for _, arg := range c.Arguments {
		_, value, err := arg.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, &Null{}, err
		}
		str += fmt.Sprint(value)
	}
	return vars, str, nil
}
