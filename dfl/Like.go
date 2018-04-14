package dfl

import (
	"fmt"
	"regexp"
	"strings"
)

import (
	"github.com/pkg/errors"
)

type Like struct {
	*BinaryOperator
}

func (l Like) Dfl() string {
	return "(" + l.Left.Dfl() + " like " + l.Right.Dfl() + ")"
}

func (l Like) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "like",
		"left":  l.Left.Map(),
		"right": l.Right.Map(),
	}
}

func (l Like) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {
	lv, err := l.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	lvs := fmt.Sprint(lv)

	rv, err := l.Right.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	rvs := fmt.Sprint(rv)

	if len(rvs) == 0 {
		return len(lvs) == 0, nil
	}

	pattern, err := regexp.Compile(strings.Replace(rvs, "%", ".*", -1))
	if err != nil {
		return false, errors.Wrap(err, "Error evaulating expression "+l.Dfl())
	}
	return pattern.MatchString(lvs), nil
}
