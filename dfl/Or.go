package dfl

import (
	"github.com/pkg/errors"
)

type Or struct {
	*BinaryOperator
}

func (o Or) Dfl() string {
	return "(" + o.Left.Dfl() + " or " + o.Right.Dfl() + ")"
}

func (o Or) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "or",
		"left":  o.Left.Map(),
		"right": o.Right.Map(),
	}
}

func (o Or) Compile() Node {
	left := o.Left.Compile()
	right := o.Right.Compile()
	return Or{&BinaryOperator{Left: left, Right: right}}
}

func (o Or) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {
	lv, err := o.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	switch lv.(type) {
	case bool:
		if lv.(bool) {
			return true, nil
		}
		rv, err := o.Right.Evaluate(ctx, funcs)
		if err != nil {
			return false, err
		}
		switch rv.(type) {
		case bool:
			return rv.(bool), nil
		}
	}
	return false, errors.New("Error evaulating expression " + o.Dfl())
}
