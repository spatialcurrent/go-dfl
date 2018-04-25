package dfl

import (
	"fmt"
	"regexp"
	"strings"
)

import (
	"github.com/pkg/errors"
)

type ILike struct {
	*BinaryOperator
}

func (i ILike) Dfl() string {
	return "(" + i.Left.Dfl() + " ilike " + i.Right.Dfl() + ")"
}

func (i ILike) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "ilike",
		"left":  i.Left.Map(),
		"right": i.Right.Map(),
	}
}

func (i ILike) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {
	lv, err := i.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	lvs := strings.ToLower(fmt.Sprint(lv))

	rv, err := i.Right.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	rvs := strings.ToLower(fmt.Sprint(rv))

	if len(rvs) == 0 {
		return len(lvs) == 0, nil
	}

	pattern, err := regexp.Compile("^" + strings.Replace(rvs, "%", ".*", -1) + "$")
	if err != nil {
		return false, errors.Wrap(err, "Error evaulating expression "+i.Dfl())
	}
	return pattern.MatchString(lvs), nil
}
