// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"

	"github.com/spatialcurrent/go-dfl/pkg/dfl/syntax"
)

// Variable is a Node representing the value of a temporary variable.
// Variables start with a "#" and follow with the name or full path into the object if multiple levels deep.
// For example, #a and #a.b.c.d.  You can also use a null-safe operator, e.g., #a?.b?.c?.d
type Variable struct {
	Name string
}

func (v Variable) Dfl(quotes []string, pretty bool, tabs int) string {
	return syntax.VariablePrefix + v.Name
}

func (v Variable) Sql(pretty bool, tabs int) string {
	if pretty {
		return strings.Repeat("  ", tabs) + v.Name
	}
	return v.Name
}

func (v Variable) Map() map[string]interface{} {
	return map[string]interface{}{
		"variable": syntax.VariablePrefix + v.Name,
	}
}

func (v Variable) Compile() Node {
	return Variable{Name: v.Name}
}

func (v Variable) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	if len(v.Name) == 0 {
		return vars, vars, nil
	}
	value, err := Extract(v.Name, vars, vars, ctx, funcs, quotes)
	return vars, value, err
}

func (v Variable) Attributes() []string {
	return []string{}
}

func (v Variable) Variables() []string {
	return []string{v.Name}
}
