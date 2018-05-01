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
)

// Literal is a Node representing a literal/static value regardless of the context.
// The value may be of any type; however, it will likely a bool, int, or string.
// For example
//	Literal{Value: ""}
//	Literal{Value: 0.0}
type Literal struct {
	Value interface{} // the variable containing the actual value
}

func (l Literal) Dfl() string {
	switch l.Value.(type) {
	case string:
		return fmt.Sprintf("%q", l.Value)
	case []string:
		out, _ := json.Marshal(l.Value)
		return string(out)
	}
	return fmt.Sprint(l.Value)
}

func (l Literal) Map() map[string]interface{} {
	return map[string]interface{}{
		"value": l.Value,
	}
}

func (l Literal) Compile() Node {
	return l
}

func (l Literal) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {
	return l.Value, nil
}

func (l Literal) Attributes() []string {
	return []string{}
}
