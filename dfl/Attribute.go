// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

// Attribute is a Node representing the value of an attribute in the context map.
// Attributes start with a "@" and follow with the name or full path into the object if multiple levels deep.
// For example, @a and @a.b.c.d.  You can also use a null-safe operator, e.g., @a?.b?.c?.d
type Attribute struct {
	Name string
}

func (a Attribute) Dfl(quotes []string, pretty bool, tabs int) string {
	if pretty {
		return strings.Repeat("  ", tabs) + AttributePrefix + a.Name
	}
	return AttributePrefix + a.Name
}

func (a Attribute) Sql(pretty bool, tabs int) string {
	if pretty {
		return strings.Repeat("  ", tabs) + a.Name
	}
	return a.Name
}

func (a Attribute) Map() map[string]interface{} {
	return map[string]interface{}{
		"attribute": AttributePrefix + a.Name,
	}
}

func (a Attribute) Compile() Node {
	return Attribute{Name: a.Name}
}

func (a Attribute) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	if len(a.Name) == 0 {
		return vars, ctx, nil
	}
	ctx, err := Extract(a.Name, ctx, vars, ctx, funcs, quotes)
	return vars, ctx, err
}

func (a Attribute) Attributes() []string {
	return []string{a.Name}
}
