// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package syntax

import (
	"strings"
	"unicode"
)

// IsVariable returns true if the string is a formatted variable.
func IsVariable(s string) bool {
	return strings.HasPrefix(strings.TrimLeftFunc(s, unicode.IsSpace), VariablePrefix)
}
