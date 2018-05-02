// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
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

// AddNumbers adds 2 numbers and returns the result.
// The parameters can be an int, int64, or float64.
// The parameters will be cast as applicable.
// For example you can add two integers with
//	total := AddNumbers(1, 2)
// or you could add an int with a float64.
//	total := AddNumbers(1.54345345, 5)
func AddNumbers(a interface{}, b interface{}) (interface{}, error) {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return a.(int) + b.(int), nil
		case int64:
			return int64(a.(int)) + b.(int64), nil
		case float64:
			return float64(a.(int)) + b.(float64), nil
		}
	case int64:
		switch b.(type) {
		case int:
			return a.(int64) + int64(b.(int)), nil
		case int64:
			return a.(int64) + b.(int64), nil
		case float64:
			return float64(a.(int64)) + b.(float64), nil
		}
	case float64:
		switch b.(type) {
		case int:
			return a.(float64) + float64(b.(int)), nil
		case int64:
			return a.(float64) + float64(b.(int64)), nil
		case float64:
			return a.(float64) + b.(float64), nil
		}
	}

	return 0, errors.New(fmt.Sprintf("Error adding values %#v (%v) and %#v (%v)", a, reflect.TypeOf(a).String(), b, reflect.TypeOf(b).String()))
}
