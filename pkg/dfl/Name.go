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

func Name(node Node) string {
	t := reflect.TypeOf(node)
	k := t.Kind()
	if k == reflect.Ptr {
		return fmt.Sprintf("*%s", t.Elem().Name())
	}
	return t.Name()
}
