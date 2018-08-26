// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
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

func (l Like) Dfl(quotes []string, pretty bool) string {
	return "(" + l.Left.Dfl(quotes, pretty) + " like " + l.Right.Dfl(quotes, pretty) + ")"
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

func (l Like) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {
	lv, err := l.Left.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}

	rv, err := l.Right.Evaluate(ctx, funcs, quotes)
	if err != nil {
		return false, err
	}

	match, err := CompareStrings(fmt.Sprint(lv), fmt.Sprint(rv))
	if err != nil {
		return false, errors.Wrap(err, "Error evaluating expression "+l.Dfl(quotes, false))
	}

	return match, nil
}
