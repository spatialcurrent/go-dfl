package dfl

import (
	"strconv"
	"strings"
)

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

	left_f64, err := strconv.ParseFloat(s, 64)
	if err == nil {
		return left_f64
	}

	left_int, err := strconv.Atoi(s)
	if err == nil {
		return left_int
	}

	return s
}
