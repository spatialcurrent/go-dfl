// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	//"reflect"
	"strconv"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// In is a BinaryOperator that evaluates to true if the left value is in the right value.
// The left value is cast as a string using "fmt.Sprint(lv)".
// If the right value is an array/slice, then evaluated to true if the left value is in the array/slice.
// Otherwise, evaluates to true if the right string is contained by the left string.
type In struct {
	*BinaryOperator
}

func (i In) Dfl() string {
	return "(" + i.Left.Dfl() + " in " + i.Right.Dfl() + ")"
}

func (i In) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "in",
		"left":  i.Left.Map(),
		"right": i.Right.Map(),
	}
}

func (i In) Compile() Node {
	left := i.Left.Compile()
	right := i.Right.Compile()
	return In{&BinaryOperator{Left: left, Right: right}}
}

func (i In) Evaluate(ctx Context, funcs FunctionMap) (interface{}, error) {
	lv, err := i.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, errors.Wrap(err, "Error evaulating expression "+i.Dfl())
	}
	lvs := fmt.Sprint(lv)

	rv, err := i.Right.Evaluate(ctx, funcs)
	if err != nil {
		return false, errors.Wrap(err, "Error evaulating expression "+i.Dfl())
	}

	switch rv.(type) {
	case string:
		return strings.Contains(rv.(string), lvs), nil
	case int:
		return strings.Contains(fmt.Sprint(rv), lvs), nil
	case float64:
		return strings.Contains(strconv.FormatFloat(rv.(float64), 'f', 6, 64), lvs), nil
	case []interface{}:
		for _, x := range rv.([]interface{}) {
			if lvs == fmt.Sprint(x) {
				return true, nil
			}
		}
		return false, nil
	}

	return false, errors.New("Error evaluating expression " + i.Dfl())
}
