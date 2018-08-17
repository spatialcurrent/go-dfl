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

// Last returns the last value of the elements of the given slice.  If values is not a slice, then returns an error.
func Last(values interface{}) (interface{}, error) {

	switch v := values.(type) {
	case []interface{}:
		if len(v) == 0 {
			return Null{}, nil
		}
		return v[len(v)-1], nil
	case string:
		if len(v) == 0 {
			return Null{}, nil
		}
		return v[len(v)-1], nil
	case []byte:
		if len(v) == 0 {
			return Null{}, nil
		}
		return v[len(v)-1], nil
	case []int:
		if len(v) == 0 {
			return Null{}, nil
		}
		return v[len(v)-1], nil
	case []float64:
		if len(v) == 0 {
			return Null{}, nil
		}
		return v[len(v)-1], nil
	}

	return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(values).String())
}
