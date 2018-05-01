// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Array is a Node representing the value of an attribute in the context map.
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
	return a
}

func (a Attribute) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {
	if v, ok := ctx[a.Name]; ok {
		return v, nil
	} else {
		return "", nil
	}
}

func (a Attribute) Attributes() []string {
	return []string{a.Name}
}
