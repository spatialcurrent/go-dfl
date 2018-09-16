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

func (i ILike) Dfl(quotes []string, pretty bool, tabs int) string {
	return i.BinaryOperator.Dfl("ilike", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (i ILike) Sql(pretty bool, tabs int) string {
	return i.BinaryOperator.Sql("ILIKE", pretty, tabs)
}

func (i ILike) Map() map[string]interface{} {
	return i.BinaryOperator.Map("ilike", i.Left, i.Right)
}

func (i ILike) Compile() Node {
	left := i.Left.Compile()
	right := i.Right.Compile()
	return ILike{&BinaryOperator{Left: left, Right: right}}
}

func (i ILike) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, lv, err := i.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}
	lvs := strings.ToLower(fmt.Sprint(lv))

	vars, rv, err := i.Right.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, err
	}
	rvs := strings.ToLower(fmt.Sprint(rv))

	if len(rvs) == 0 {
		return vars, len(lvs) == 0, nil
	}

	pattern, err := regexp.Compile("^" + strings.Replace(rvs, "%", ".*", -1) + "$")
	if err != nil {
		return vars, false, errors.Wrap(err, ErrorEvaluate{Node: i, Quotes: quotes}.Error())
	}
	return vars, pattern.MatchString(lvs), nil
}
