// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package syntax

import (
	"strings"
	"unicode"
)

// IsAttribute returns true if the string is a formatted attribute
func IsAttribute(s string) bool {
	return strings.HasPrefix(strings.TrimLeftFunc(s, unicode.IsSpace), AttributePrefix)
}
