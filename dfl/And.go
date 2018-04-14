package dfl

import (
	"github.com/pkg/errors"
)

type And struct {
	*BinaryOperator
}

func (a And) Dfl() string {
	return "(" + a.Left.Dfl() + " and " + a.Right.Dfl() + ")"
}

func (a And) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "and",
		"left":  a.Left.Map(),
		"right": a.Right.Map(),
	}
}

func (a And) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {
	lv, err := a.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	switch lv.(type) {
	case bool:
		if !lv.(bool) {
			return false, nil
		}
		rv, err := a.Right.Evaluate(ctx, funcs)
		if err != nil {
			return false, err
		}
		switch rv.(type) {
		case bool:
			return rv.(bool), nil
		}
	}
	return false, errors.New("Error evaulating expression " + a.Dfl())
}
