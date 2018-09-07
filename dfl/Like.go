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

func (l Like) Dfl(quotes []string, pretty bool, tabs int) string {
	return l.BinaryOperator.Dfl("like", quotes, pretty, tabs)
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

func (l Like) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, lv, err := l.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}

	vars, rv, err := l.Right.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}

	match, err := CompareStrings(fmt.Sprint(lv), fmt.Sprint(rv))
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating expression "+l.Dfl(quotes, false, 0))
	}

	return vars, match, nil
}
