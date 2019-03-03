// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

type Lengther interface {
	Len() int
}

func CreateGroups(depth int) interface{} {

	switch depth {
	case 1:
		return map[string][]interface{}{}
	case 2:
		return map[string]map[string][]interface{}{}
	case 3:
		return map[string]map[string]map[string][]interface{}{}
	case 4:
		return map[string]map[string]map[string]map[string][]interface{}{}
	case 5:
		return map[string]map[string]map[string]map[string]map[string][]interface{}{}
	case 6:
		return map[string]map[string]map[string]map[string]map[string]map[string][]interface{}{}
	case 7:
		return map[string]map[string]map[string]map[string]map[string]map[string]map[string][]interface{}{}
	case 8:
		return map[string]map[string]map[string]map[string]map[string]map[string]map[string]map[string][]interface{}{}
	}

	return nil
}
