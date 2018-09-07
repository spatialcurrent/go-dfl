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

// Variable is a Node representing the value of a temporary variable.
// Variables start with a "#" and follow with the name or full path into the object if multiple levels deep.
// For example, #a and #a.b.c.d.  You can also use a null-safe operator, e.g., #a?.b?.c?.d
type Variable struct {
	Name string
}

func (v Variable) Dfl(quotes []string, pretty bool, tabs int) string {
	if pretty {
		return strings.Repeat("  ", tabs) + VariablePrefix + v.Name
	}
	return VariablePrefix + v.Name
}

func (v Variable) Map() map[string]interface{} {
	return map[string]interface{}{
		"attribute": VariablePrefix + v.Name,
	}
}

func (v Variable) Compile() Node {
	return Variable{Name: v.Name}
}

func (v Variable) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	if len(v.Name) == 0 {
		return vars, ctx, errors.New("variable")
	}
	value, err := Extract(v.Name, vars, vars, ctx, funcs, quotes)
	return vars, value, err
}

func (v Variable) Attributes() []string {
	return []string{}
}
