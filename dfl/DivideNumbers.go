// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"reflect"
)

import (
	"github.com/pkg/errors"
)

// DivideNumbers divides the first number by the second number and returns the results.
// The parameters can be an int, int64, or float64.
// The parameters will be cast as applicable.
// For example you can divide two integers with
//	total := DivideNumbers(1, 2)
// or you could divide an int with a float64.
//	total := DivideNumbers(4, 3.2)
func DivideNumbers(a interface{}, b interface{}) (interface{}, error) {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return a.(int) / b.(int), nil
		case int64:
			return int64(a.(int)) / b.(int64), nil
		case float64:
			return float64(a.(int)) / b.(float64), nil
		}
	case int64:
		switch b.(type) {
		case int:
			return a.(int64) / int64(b.(int)), nil
		case int64:
			return a.(int64) / b.(int64), nil
		case float64:
			return float64(a.(int64)) / b.(float64), nil
		}
	case float64:
		switch b.(type) {
		case int:
			return a.(float64) / float64(b.(int)), nil
		case int64:
			return a.(float64) / float64(b.(int64)), nil
		case float64:
			return a.(float64) / b.(float64), nil
		}
	}

	return 0, errors.New(fmt.Sprintf("Error dividing %#v (%v) - %#v (%v)", a, reflect.TypeOf(a).String(), b, reflect.TypeOf(b).String()))
}
