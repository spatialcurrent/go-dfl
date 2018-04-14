package dfl

import (
	"fmt"
	"strings"
)

import (
	"github.com/pkg/errors"
)

type In struct {
	*BinaryOperator
}

func (i In) Dfl() string {
	return "(" + i.Left.Dfl() + " in " + i.Right.Dfl() + ")"
}

func (i In) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "in",
		"left":  i.Left.Map(),
		"right": i.Right.Map(),
	}
}

func (i In) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {
	lv, err := i.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, errors.Wrap(err, "Error evaulating expression "+i.Dfl())
	}
	lvs := fmt.Sprint(lv)

	rv, err := i.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, errors.Wrap(err, "Error evaulating expression "+i.Dfl())
	}
	rvs := fmt.Sprint(rv)

	return strings.Contains(lvs, rvs), nil
}
