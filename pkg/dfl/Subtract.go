// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
)

type Subtract struct {
	*NumericBinaryOperator
}

func (s Subtract) Dfl(quotes []string, pretty bool, tabs int) string {
	return s.BinaryOperator.Dfl("-", quotes, pretty, tabs)
}

func (s Subtract) Sql(pretty bool, tabs int) string {
	return s.BinaryOperator.Sql("-", pretty, tabs)
}

func (s Subtract) Map() map[string]interface{} {
	return s.BinaryOperator.Map("subtract", s.Left, s.Right)
}

func (s Subtract) Compile() Node {
	left := s.Left.Compile()
	right := s.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := af.Subtract.ValidateRun([]interface{}{left.(Literal).Value, right.(Literal).Value})
			if err != nil {
				return &Subtract{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
			}
			return Literal{Value: v}
		}
	}
	return &Subtract{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (s Subtract) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, rv, err := s.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	v, err := af.Subtract.ValidateRun(lv, rv)
	if err != nil {
		return vars, 0, errors.Wrap(err, ErrorEvaluate{Node: s, Quotes: quotes}.Error())
	}

	return vars, v, nil
}
