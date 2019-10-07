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
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
)

// In is a BinaryOperator that evaluates to true if the left value is in the right value.
// The left value is cast as a string using "fmt.Sprint(lv)".
// If the right value is an array/slice, then evaluated to true if the left value is in the array/slice.
// Otherwise, evaluates to true if the right string is contained by the left string.
type In struct {
	*BinaryOperator
}

func (i In) Dfl(quotes []string, pretty bool, tabs int) string {
	return i.BinaryOperator.Dfl("in", quotes, pretty, tabs)
}

func (i In) Sql(pretty bool, tabs int) string {

	switch right := i.Right.(type) {
	case *Attribute:
		switch left := i.Left.(type) {
		case *Literal:
			switch lv := left.Value.(type) {
			case string:
				like := &Like{&BinaryOperator{
					Left:  i.Right,
					Right: &Literal{Value: "%" + lv + "%"},
				}}
				return like.Sql(pretty, tabs)
			}
		}
	case *Set:
		eq := &Equal{&BinaryOperator{
			Left:  i.Left,
			Right: &Function{Name: "ANY", MultiOperator: &MultiOperator{Arguments: []Node{right}}},
		}}
		return eq.Sql(pretty, tabs)
	}

	return ""
}

func (i In) Map() map[string]interface{} {
	return i.BinaryOperator.Map("in", i.Left, i.Right)
}

func (i In) Compile() Node {
	left := i.Left.Compile()
	right := i.Right.Compile()
	return &In{&BinaryOperator{Left: left, Right: right}}
}

func (i In) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, err := i.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating left value for "+i.Dfl(quotes, false, 0))
	}

	vars, rv, err := i.Right.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating right value for "+i.Dfl(quotes, false, 0))
	}

	if rvr, ok := rv.(io.ByteReadCloser); ok {
		if lvb, ok := lv.([]byte); ok {
			rvb, err := rvr.ReadAll()
			if err != nil {
				return vars, false, errors.Wrap(err, "error reading all byte for right value in expression "+i.Dfl(quotes, false, 0))
			}
			if len(lvb) == len(rvb) && len(lvb) == 0 {
				return vars, true, nil
			}
			for i, _ := range rvb {
				if rvb[i] == lvb[0] && i+len(lvb) < len(rvb) {
					match := true
					for j, _ := range lvb {
						if rvb[i+j] != lvb[j] {
							match = false
							break
						}
					}
					if match {
						return vars, true, nil
					}
				}
			}
			return vars, false, nil
		}
		if lvs, ok := lv.(string); ok {
			lvb := []byte(lvs)
			rvb, err := rvr.ReadAll()
			if err != nil {
				return vars, false, errors.Wrap(err, "error reading all byte for right value in expression "+i.Dfl(quotes, false, 0))
			}
			if len(lvb) == len(rvb) && len(lvb) == 0 {
				return vars, true, nil
			}
			for i, _ := range rvb {
				if rvb[i] == lvb[0] && i+len(lvb) < len(rvb) {
					match := true
					for j, _ := range lvb {
						if rvb[i+j] != lvb[j] {
							match = false
							break
						}
					}
					if match {
						return vars, true, nil
					}
				}
			}
			return vars, false, nil
		}
	}

	value, err := af.In.ValidateRun(lv, rv)
	if err != nil {
		return vars, false, errors.Wrap(err, ErrorEvaluate{Node: i, Quotes: quotes}.Error())
	}

	return vars, value, nil

}
