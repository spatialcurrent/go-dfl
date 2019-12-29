// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package syntax

// IsNull returns true if the string is a formatted null.
func IsNull(s string) bool {
	return s == "null" || s == "none" || s == "nil"
}
