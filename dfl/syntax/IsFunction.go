// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package syntax

import (
	"strings"
)

// IsFunction returns true if the string is a formatted function
func IsFunction(s string) bool {
	return len(s) >= 3 && (strings.Index(s, "(") > 0) && strings.HasSuffix(s, ")")
}
