// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

// IsFunction returns true if the string is a formatted function
func IsFunction(s string) bool {
	return len(s) >= 3 && strings.Contains(s, "(") && strings.HasSuffix(s, ")")
}
