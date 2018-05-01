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

// SubtractNumbers subtracts the second number from the first and returns the results.
// The parameters can be an int, int64, or float64.
// The parameters will be cast as applicable.
// For example you can add two integers with
//	total := SubtractNumbers(1, 2)
// or you could subtract an int with a float64.
//	total := AddNumbers(4, 3.2)
func SubtractNumbers(a interface{}, b interface{}) (interface{}, error) {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return a.(int) - b.(int), nil
		case int64:
			return int64(a.(int)) - b.(int64), nil
		case float64:
			return float64(a.(int)) - b.(float64), nil
		}
	case int64:
		switch b.(type) {
		case int:
			return a.(int64) - int64(b.(int)), nil
		case int64:
			return a.(int64) - b.(int64), nil
		case float64:
			return float64(a.(int64)) - b.(float64), nil
		}
	case float64:
		switch b.(type) {
		case int:
			return a.(float64) - float64(b.(int)), nil
		case int64:
			return a.(float64) - float64(b.(int64)), nil
		case float64:
			return a.(float64) - b.(float64), nil
		}
	}

	return 0, errors.New("Error subtracting " + fmt.Sprint(a) + " - " + fmt.Sprint(b))
}
