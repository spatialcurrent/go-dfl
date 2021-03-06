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

	"github.com/pkg/errors"
)

// AttachLeft attaches the left Node as the left child node to the parent root Node.
func AttachLeft(root Node, left Node) error {

	t := reflect.TypeOf(root)
	v := reflect.ValueOf(root)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		return errors.New("could not attach left as root " + fmt.Sprint(t) + " is not a struct but " + fmt.Sprint(t))
	}
	f := v.FieldByName("Left")
	if !f.IsValid() {
		return errors.New("could not attach left as root " + fmt.Sprint(t) + " does not have a field with name Left")
	}
	if !f.CanSet() {
		return errors.New("could not attach left as root " + fmt.Sprint(t) + " does not have a field with name Left that can be set")
	}
	f.Set(reflect.ValueOf(left))

	return nil
}
