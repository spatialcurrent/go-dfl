// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"reflect"
)

func TryFormatLiteral(value interface{}, quotes []string, pretty bool, tabs int) string {

	t := reflect.TypeOf(value)
	if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		v := reflect.ValueOf(value)
		s := "["
		l := v.Len()
		if l > 0 {
			for i := 0; i < l; i++ {
				s += TryFormatLiteral(v.Index(i).Interface(), quotes, false, 0)
				if i < l-1 {
					s += ", "
				}
			}
		}
		s += "]"
		return s
	}

	switch value := value.(type) {
	case string:
		return quotes[0] + value + quotes[0]
	case map[string]struct{}:
		return StringSet(value).Dfl(quotes, pretty, tabs)
	case StringSet:
		return value.Dfl(quotes, pretty, tabs)
	case Null:
		return value.Dfl()
	}

	return fmt.Sprint(value)
}
