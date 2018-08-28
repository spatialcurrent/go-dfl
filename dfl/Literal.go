// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Literal is a Node representing a literal/static value regardless of the context.
// The value may be of any type; however, it will likely a bool, int, or string.
// For example
//	Literal{Value: ""}
//	Literal{Value: 0.0}
type Literal struct {
	Value interface{} // the variable containing the actual value
}

func (l Literal) Dfl(quotes []string, pretty bool, tabs int) string {

	str := ""
	switch value := l.Value.(type) {
	case string:
		str = quotes[0] + value + quotes[0]
		//return fmt.Sprintf("%q", value)
	case []string:
		out, _ := json.Marshal(value)
		str = string(out)
	case map[string]struct{}:
		str = StringSet(value).Dfl(quotes, pretty, tabs)
	case StringSet:
		str = value.Dfl(quotes, pretty, tabs)
	case Null:
		str = value.Dfl()
	default:
		str = fmt.Sprint(l.Value)
	}

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

func (l Literal) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {
	return l.Value, nil
}

func (l Literal) Attributes() []string {
	return []string{}
}
