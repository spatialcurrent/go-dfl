// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Function is a refrenced function in a DFL filter.  The actual function in a given FunctionMap is derefernced by name.
type Function struct {
	Name      string `json:"name" bson:"name" yaml:"name" hcl:"name"`                     // name of the function
	Arguments []Node `json:"arguments" bson:"arguments" yaml:"arguments" hcl:"arguments"` // list of function arguments
}

func (f Function) Dfl() string {
	out := f.Name + "("
	for i, arg := range f.Arguments {
		if i > 0 {
			out += ", "
		}
		out += arg.Dfl()
	}
	out += ")"
	return out
}

func (f Function) Compile() Node {
	return f
}

func (f Function) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":        "function",
		"name":      f.Name,
		"arguments": f.Arguments,
	}
}

func (f Function) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {
	if fn, ok := funcs[f.Name]; ok {
		values := make([]interface{}, 0, len(f.Arguments))
		for _, arg := range f.Arguments {
			value, err := arg.Evaluate(ctx, funcs)
			if err != nil {
				return &Null{}, err
			}
			values = append(values, value)
		}
		return fn(ctx, values)
	} else {
		return "", nil
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
