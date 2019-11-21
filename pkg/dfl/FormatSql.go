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
)

// FormatSQL formats an object as a SQL string
func FormatSql(value interface{}, pretty bool, tabs int) string {

	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		v := reflect.ValueOf(value)
		s := SqlArrayPrefix
		l := v.Len()
		if l > 0 {
			for i := 0; i < l; i++ {
				s += FormatSql(v.Index(i).Interface(), false, 0)
				if i < l-1 {
					s += ", "
				}
			}
		}
		s += SqlArraySuffix
		return s
	}

	switch value := value.(type) {
	case string:
		return SqlQuote + value + SqlQuote
	case map[string]struct{}:
		// Just format sets as slices.
		return FormatSql(StringSet(value).Slice(true), pretty, tabs)
	case StringSet:
		// Just format sets as slices.
		return FormatSql(value.Slice(true), pretty, tabs)
	case Null:
		return value.Sql(pretty, tabs)
	}

	return fmt.Sprint(value)
}
