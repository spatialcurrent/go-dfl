// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
)

import (
	"github.com/pkg/errors"
)

// Max returns the maximum value of the elements of the given slice.  If values is not a slice, then returns an error.
func Max(values interface{}) (interface{}, error) {
	switch v := values.(type) {
	case []int:
		if len(v) == 0 {
			return Null{}, errors.New("Invalid length of " + reflect.TypeOf(v).String())
		}
		max := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] > max {
				max = v[i]
			}
		}
		return max, nil
	case []float64:
		if len(v) == 0 {
			return Null{}, errors.New("Invalid length of " + reflect.TypeOf(v).String())
		}
		max := v[0]
		for i := 1; i < len(v); i++ {
			if v[i] > max {
				max = v[i]
			}
		}
		return max, nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(values).String())
}
