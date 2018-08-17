// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// ILike is a BinaryOperator that evaluates the SQL standard like expression.
// It is similar to the Like BinaryOperator but is case insensitive.
// The parameters must be of type string.
// The right parameter may have "%" characters that are interpreted as (.*) in a regular expression test.
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

func (i ILike) Compile() Node {
	left := i.Left.Compile()
	right := i.Right.Compile()
	return ILike{&BinaryOperator{Left: left, Right: right}}
}

func (i ILike) Evaluate(ctx interface{}, funcs FunctionMap) (interface{}, error) {
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
		return false, errors.Wrap(err, "Error evaluating expression "+i.Dfl())
	}
	return pattern.MatchString(lvs), nil
}
