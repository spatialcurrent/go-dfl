// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-adaptive-functions/af"
)

// In is a BinaryOperator that evaluates to true if the left value is in the right value.
// Unlike "in", it is case insensitive.
// If the right value is an array/slice, then evaluated to true if the left value is in the array/slice.
// Otherwise, evaluates to true if the right string is contained by the left string.
type IIn struct {
	*BinaryOperator
}

func (i IIn) Dfl(quotes []string, pretty bool, tabs int) string {
	return i.BinaryOperator.Dfl("iin", quotes, pretty, tabs)
}

func (i IIn) Sql(pretty bool, tabs int) string {

	return ""
}

func (i IIn) Map() map[string]interface{} {
	return i.BinaryOperator.Map("iin", i.Left, i.Right)
}

func (i IIn) Compile() Node {
	left := i.Left.Compile()
	right := i.Right.Compile()
	return &IIn{&BinaryOperator{Left: left, Right: right}}
}

func (i IIn) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, err := i.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating left value for "+i.Dfl(quotes, false, 0))
	}

	vars, rv, err := i.Right.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating right value for "+i.Dfl(quotes, false, 0))
	}

	value, err := af.IIn.ValidateRun([]interface{}{lv, rv})
	if err != nil {
		return vars, false, errors.Wrap(err, ErrorEvaluate{Node: i, Quotes: quotes}.Error())
	}

	return vars, value, nil

}
