// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Null is used as return value for Extract and DFL functions instead of returning nil pointers.
type Null struct{}

func (n Null) Dfl() string {
	return "null"
}

func (n Null) Sql() string {
	return "NULL"
}
