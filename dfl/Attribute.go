// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Attribute is a Node representing the value of an attribute in the context map.
// Attributes start with a "@" and follow with the name or full path into the object if multiple levels deep.
// For example, @a and @a.b.c.d.  You can also use a null-safe operator, e.g., @a?.b?.c?.d
type Attribute struct {
	Name string
}

func (a Attribute) Dfl() string {
	return "@" + a.Name
}

func (a Attribute) Map() map[string]interface{} {
	return map[string]interface{}{
		"attribute": a.Name,
	}
}

func (a Attribute) Compile() Node {
	return Attribute{Name: a.Name}
}

func (a Attribute) Evaluate(ctx interface{}, funcs FunctionMap) (interface{}, error) {
	return Extract(a.Name, ctx)
}

func (a Attribute) Attributes() []string {
	return []string{a.Name}
}
