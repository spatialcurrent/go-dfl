// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package syntax

import (
	"strings"
)

// IsSub returns true if the string is formatted sub
func IsSub(s string) bool {
	return len(s) >= 2 && strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")")
}
