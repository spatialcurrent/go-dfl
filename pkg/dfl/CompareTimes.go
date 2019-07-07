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
	"time"
)

import (
	"github.com/pkg/errors"
)

// CompareTimes compares parameter a and parameter b.
// The parameters may be of type string, time.Time, or *time.Time.
// If a is after b, then returns 1.  If a is before b, then returns -1.  If a is at the same time as b, then return 0.
func CompareTimes(a interface{}, b interface{}) (int, error) {
	switch a.(type) {
	case string:
		at, err := time.Parse(time.RFC3339, a.(string))
		if err != nil {
			return 0, errors.New("Error parsing value " + (a.(string)))
		}
		switch b.(type) {
		case string:
			bt, err := time.Parse(time.RFC3339, b.(string))
			if err != nil {
				return 0, errors.New("Error parsing value " + (b.(string)))
			}
			if at.Before(bt) {
				return -1, nil
			} else if at.After(bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		case time.Time:
			bt := b.(time.Time)
			if at.Before(bt) {
				return -1, nil
			} else if at.After(bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		case *time.Time:
			bt := b.(*time.Time)
			if at.Before(*bt) {
				return -1, nil
			} else if at.After(*bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case time.Time:
		at := a.(time.Time)
		switch b.(type) {
		case string:
			bt, err := time.Parse(time.RFC3339, b.(string))
			if err != nil {
				return 0, errors.New("Error parsing value " + (b.(string)))
			}
			if at.Before(bt) {
				return -1, nil
			} else if at.After(bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		case time.Time:
			bt := b.(time.Time)
			if at.Before(bt) {
				return -1, nil
			} else if at.After(bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		case *time.Time:
			bt := b.(*time.Time)
			if at.Before(*bt) {
				return -1, nil
			} else if at.After(*bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	case *time.Time:
		at := a.(*time.Time)
		switch b.(type) {
		case string:
			bt, err := time.Parse(time.RFC3339, b.(string))
			if err != nil {
				return 0, errors.New("Error parsing value " + (b.(string)))
			}
			if (*at).Before(bt) {
				return -1, nil
			} else if (*at).After(bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		case time.Time:
			bt := b.(time.Time)
			if (*at).Before(bt) {
				return -1, nil
			} else if (*at).After(bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		case *time.Time:
			bt := b.(*time.Time)
			if (*at).Before(*bt) {
				return -1, nil
			} else if (*at).After(*bt) {
				return 1, nil
			} else {
				return 0, nil
			}
		}
	}
	return 0, errors.New(fmt.Sprintf("Error comparing times %#v (%v) and %#v (%v)", a, reflect.TypeOf(a).String(), b, reflect.TypeOf(b).String()))
}
