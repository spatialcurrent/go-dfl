// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
)

// WithinRange returns true if value is in the range [start, end]
func WithinRange(value interface{}, start interface{}, end interface{}) (bool, error) {
	v, err := CompareNumbers(value, start)
	if err != nil {
		return false, errors.Wrap(err, "error in WithinRange comparing value to start")
	}
	if v < 0 {
		return false, nil
	}
	v, err = CompareNumbers(value, end)
	if err != nil {
		return false, errors.Wrap(err, "error in WithinRange comparing value to end")
	}
	return v <= 0, nil
}
