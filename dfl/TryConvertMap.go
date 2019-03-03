// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
)

// TryConvertMap attempts to convert the interface{} map into a typed version
func TryConvertMap(m interface{}) interface{} {

	mt := reflect.TypeOf(m)
	if mt.Kind() != reflect.Map {
		return m
	}

	mv := reflect.ValueOf(m)
	if mv.Len() == 0 {
		return m
	}

	keyTypes := map[reflect.Type]struct{}{}
	valueTypes := map[reflect.Type]struct{}{}

	for _, k := range mv.MapKeys() {
		kt := reflect.TypeOf(k.Interface())
		if _, ok := keyTypes[kt]; !ok {
			keyTypes[kt] = struct{}{}
		}
		v := mv.MapIndex(k)
		vt := reflect.TypeOf(v.Interface())
		if _, ok := valueTypes[vt]; !ok {
			valueTypes[vt] = struct{}{}
		}
	}

	if len(keyTypes) == 1 || len(valueTypes) == 1 {
		keyType := mt.Key()
		if len(keyTypes) == 1 {
			for x, _ := range keyTypes {
				keyType = x
			}
		}

		valueType := mt.Elem()
		if len(valueTypes) == 1 {
			for x, _ := range valueTypes {
				valueType = x
			}
		}

		n := reflect.MakeMap(reflect.MapOf(keyType, valueType))
		for _, k := range mv.MapKeys() {
			n.SetMapIndex(reflect.ValueOf(k.Interface()), reflect.ValueOf(mv.MapIndex(k).Interface()))
		}

		return n.Interface()
	}

	return m
}
