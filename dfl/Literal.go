// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

// Literal is a Node representing a literal/static value regardless of the context.
// The value may be of any type; however, it will likely a bool, int, or string.
// For example
//	Literal{Value: ""}
//	Literal{Value: 0.0}
type Literal struct {
	Value interface{} // the field containing the actual value
}

func (l Literal) Dfl(quotes []string, pretty bool, tabs int) string {
	return TryFormatLiteral(l.Value, quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (l Literal) Sql(pretty bool, tabs int) string {
	str := FormatSql(l.Value, pretty, tabs)

	if pretty {
		str = strings.Repeat("  ", tabs) + str
	}

	return str
}

func (l Literal) Map() map[string]interface{} {
	return map[string]interface{}{
		"value": l.Value,
	}
}

func (l Literal) Compile() Node {
	return l
}

func (l Literal) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	return vars, l.Value, nil
}

func (l Literal) Attributes() []string {
	return []string{}
}

func (l Literal) Variables() []string {
	return []string{}
}
