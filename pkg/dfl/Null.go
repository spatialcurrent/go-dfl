// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Null is used as return value for Extract and DFL functions instead of returning nil pointers.
type Null struct{}

func (n Null) Dfl(quotes []string, pretty bool, tabs int) string {
	return "null"
}

func (n Null) Sql(pretty bool, tabs int) string {
	return "NULL"
}

func (n Null) Map() map[string]interface{} {
	return map[string]interface{}{
		"@type": "null",
	}
}

func (n Null) MarshalMap() (interface{}, error) {
	return n.Map(), nil
}

func (n Null) Compile() Node {
	return &Null{}
}

func (n Null) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	return vars, nil, nil
}

func (n Null) Attributes() []string {
	return []string{}
}

func (n Null) Variables() []string {
	return []string{}
}
