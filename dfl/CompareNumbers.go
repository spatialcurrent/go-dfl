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

// CompareNumbers compares parameter a and parameter b.
// The parameters may be of type int, int64, or float64.
// If a > b, then returns 1.  If a < b, then returns -1.  If a == b, then return 0.
func CompareNumbers(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			if a.(int) > b.(int) {
				return 1, nil
			} else if a.(int) < b.(int) {
				return -1, nil
			} else {
				return 0, nil
			}
		case int64:
			if int64(a.(int)) > b.(int64) {
				return 1, nil
			} else if int64(a.(int)) < b.(int64) {
				return -1, nil
			} else {
				return 0, nil
			}
		case float64:
			if float64(a.(int)) > b.(float64) {
				return 1, nil
			} else if float64(a.(int)) < b.(float64) {
				return -1, nil
			} else {
				return 0, nil
			}
		}
	case int64:
		switch b.(type) {
		case int:
			if a.(int64) > int64(b.(int)) {
				return 1, nil
			} else if a.(int64) < int64(b.(int)) {
				return -1, nil
			} else {
				return 0, nil
			}
		case int64:
			if a.(int64) > b.(int64) {
				return 1, nil
			} else if a.(int64) < b.(int64) {
				return -1, nil
			} else {
				return 0, nil
			}
		case float64:
			if float64(a.(int64)) > b.(float64) {
				return 1, nil
			} else if float64(a.(int64)) < b.(float64) {
				return -1, nil
			} else {
				return 0, nil
			}
		}
	case float64:
		switch b.(type) {
		case int:
			if a.(float64) > float64(b.(int)) {
				return 1, nil
			} else if a.(float64) < float64(b.(int)) {
				return -1, nil
			} else {
				return 0, nil
			}
		case int64:
			if a.(float64) > float64(b.(int64)) {
				return 1, nil
			} else if a.(float64) < float64(b.(int64)) {
				return -1, nil
			} else {
				return 0, nil
			}
		case float64:
			if a.(float64) > b.(float64) {
				return 1, nil
			} else if a.(float64) < b.(float64) {
				return -1, nil
			} else {
				return 0, nil
			}
		}
	}

	return 0, errors.New("Error comparing values " + fmt.Sprint(a) + " and " + fmt.Sprint(b))
}
