// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"regexp"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// Like is a BinaryOperator that evaluates the SQL standard like expression.
// It is similar to the ILike BinaryOperator but is case sensitive.
// The parameters must be of type string.
// The right parameter may have "%" characters that are interpreted as (.*) in a regular expression test.
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

func (l Like) Compile() Node {
	left := l.Left.Compile()
	right := l.Right.Compile()
	return Like{&BinaryOperator{Left: left, Right: right}}
}

func (l Like) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {
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

	pattern, err := regexp.Compile("^" + strings.Replace(rvs, "%", ".*", -1) + "$")
	if err != nil {
		return false, errors.Wrap(err, "Error evaulating expression "+l.Dfl())
	}
	return pattern.MatchString(lvs), nil
}
