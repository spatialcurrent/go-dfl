// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/spatialcurrent/go-counter/counter"
	"reflect"
)

// TryConvertArray attempts to convert the []interface{} array into []int, []int64, []float64, or []string, if possible.
func TryConvertArray(a []interface{}) interface{} {
	counter := counter.New()
	for _, v := range a {
		counter.Increment(reflect.TypeOf(v).String())
	}

	switch counter.Len() {
	case 1:
		if counter.Has("int") {
			arrInt := make([]int, 0, len(a))
			for _, v := range a {
				arrInt = append(arrInt, v.(int))
			}
			return arrInt
		} else if counter.Has("int64") {
			arrInt64 := make([]int64, 0, len(a))
			for _, v := range a {
				arrInt64 = append(arrInt64, v.(int64))
			}
			return arrInt64
		} else if counter.Has("uint8") {
			arrUInt8 := make([]uint8, 0, len(a))
			for _, v := range a {
				arrUInt8 = append(arrUInt8, v.(uint8))
			}
			return arrUInt8
		} else if counter.Has("float64") {
			arrFloat64 := make([]float64, 0, len(a))
			for _, v := range a {
				arrFloat64 = append(arrFloat64, v.(float64))
			}
			return arrFloat64
		} else if counter.Has("string") {
			arrString := make([]string, 0, len(a))
			for _, v := range a {
				arrString = append(arrString, v.(string))
			}
			return arrString
		}
	case 2:
		if counter.Has("int") && counter.Has("uint8") {
			arrInt := make([]int, 0, len(a))
			for _, v := range a {
				switch v.(type) {
				case int:
					arrInt = append(arrInt, v.(int))
				case uint8:
					arrInt = append(arrInt, int(v.(uint8)))
				}
			}
			return arrInt
		} else if counter.Has("int") && counter.Has("int64") {
			arrInt64 := make([]int64, 0, len(a))
			for _, v := range a {
				switch v.(type) {
				case int:
					arrInt64 = append(arrInt64, int64(v.(int)))
				case int64:
					arrInt64 = append(arrInt64, v.(int64))
				}
			}
			return arrInt64
		} else if counter.Has("int") && counter.Has("float64") {
			arrFloat64 := make([]float64, 0, len(a))
			for _, v := range a {
				switch v.(type) {
				case int:
					arrFloat64 = append(arrFloat64, float64(v.(int)))
				case float64:
					arrFloat64 = append(arrFloat64, v.(float64))
				}
			}
			return arrFloat64
		} else if counter.Has("int64") && counter.Has("float64") {
			arrFloat64 := make([]float64, 0, len(a))
			for _, v := range a {
				switch v.(type) {
				case int64:
					arrFloat64 = append(arrFloat64, float64(v.(int64)))
				case float64:
					arrFloat64 = append(arrFloat64, v.(float64))
				}
			}
			return arrFloat64
		}
	case 3:
		if counter.Has("int") && counter.Has("int64") && counter.Has("float64") {
			arrFloat64 := make([]float64, 0, len(a))
			for _, v := range a {
				switch v.(type) {
				case int:
					arrFloat64 = append(arrFloat64, float64(v.(int)))
				case int64:
					arrFloat64 = append(arrFloat64, float64(v.(int64)))
				case float64:
					arrFloat64 = append(arrFloat64, v.(float64))
				}
			}
			return arrFloat64
		}
	}

	return a
}
