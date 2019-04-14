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
	"strings"
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
		if len(value) == 0 {
			return quotes[0] + quotes[0]
		}
		str := EscapeString(value)
		if strings.ContainsAny(str, "\\-+=,.`'\"?:[]()\n\t ") {
			return quotes[0] + str + quotes[0]
		}
		return str
	case map[string]struct{}:
		return StringSet(value).Dfl(quotes, pretty, tabs)
	case StringSet:
		return value.Dfl(quotes, pretty, tabs)
	case Null:
		return value.Dfl()
	}

	return fmt.Sprint(value)
}
