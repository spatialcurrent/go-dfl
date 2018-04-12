package dfl

import (
	"strings"
)

type Function struct {
	Name      string   `json:"name" bson:"name" yaml:"name" hcl:"name"`
	Arguments []string `json:"arguments" bson:"arguments" yaml:"arguments" hcl:"arguments"`
}

func (f Function) Dfl() string {
	out := f.Name + "("
	for i, arg := range f.Arguments {
		if i > 0 {
			out += ", "
		}
		if strings.Contains(arg, " ") {
			out += "\"" + arg + "\""
		} else {
			out += arg
		}
	}
	out += ")"
	return out
}

func (f Function) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":        "function",
		"name":      f.Name,
		"arguments": f.Arguments,
	}
}
