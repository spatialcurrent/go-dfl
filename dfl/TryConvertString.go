// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strconv"
	"strings"
	"time"
)

// TryConvertString attempts to parse the string parameter s into an appropriate literal value of type string, bool, int, float64, or time.Time.
// The functions evaluates the following rules in order.  It returns the first sucess.  The rules are:
//
//	1. "null", "none", "" => ""
//	2. "true" => true (bool)
//	3. "false" => false (bool)
//	4. "0.234" => float64
//	5. 131238 => int
//	6. time.Parse(time.RFC3339Nano, s)
//	7. time.Parse(time.RFC3339, s)
//	8. time.Parse("2006-01-02", s)
//	9. If no rules pass without error, then just return the input value
//
// For example:
//	TryConvertString("a") => "a" (string)
//	TryConvertString("true") => true (bool)
//	TryConvertString("123.31") => 123.31 (float64)
//	TryConvertString("4") => 4 (int)
//	TryConvertString("2018-05-01") => 2018-05-01T00:00:00Z (time.Time)
func TryConvertString(s string) interface{} {
	s_lc := strings.ToLower(s)

	if s_lc == "null" || s_lc == "none" || s_lc == "" {
		return ""
	}

	if s_lc == "true" {
		return true
	}

	if s_lc == "false" {
		return false
	}

	if strings.Contains(s, ".") {
		left_f64, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return left_f64
		}
	}

	left_int, err := strconv.Atoi(s)
	if err == nil {
		return left_int
	}

	left_time, err := time.Parse(time.RFC3339Nano, s)
	if err == nil {
		return left_time
	}

	left_time, err = time.Parse(time.RFC3339, s)
	if err == nil {
		return left_time
	}

	left_time, err = time.Parse("2006-01-02", s)
	if err == nil {
		return left_time
	}

	return s
}
