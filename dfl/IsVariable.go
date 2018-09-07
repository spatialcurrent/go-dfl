// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
	"unicode"
)

// IsVariable returns true if the string is a formatted variable.
func IsVariable(s string) bool {
	return strings.HasPrefix(strings.TrimLeftFunc(s, unicode.IsSpace), VariablePrefix)
}
