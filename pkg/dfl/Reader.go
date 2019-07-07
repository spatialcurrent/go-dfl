// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

type Reader interface {
	ReadAll() ([]byte, error)
	ReadRange(start int, end int) ([]byte, error)
}
