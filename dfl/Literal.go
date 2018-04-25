package dfl

import (
	"encoding/json"
	"fmt"
	//"strings"
)

type Literal struct {
	Value interface{}
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
func (l Literal) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {
	return l.Value, nil
}

func (l Literal) Attributes() []string {
	return []string{}
}
