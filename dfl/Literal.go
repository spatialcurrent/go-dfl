package dfl

import (
	"fmt"
)

type Literal struct {
	Value interface{}
}

func (l Literal) Dfl() string {
	switch l.Value.(type) {
	case string:
		return fmt.Sprintf("%q", l.Value)
	}
	return fmt.Sprint(l.Value)
}

func (l Literal) Map() map[string]interface{} {
	return map[string]interface{}{
		"value": l.Value,
	}
}
func (l Literal) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {
	return l.Value, nil
}
