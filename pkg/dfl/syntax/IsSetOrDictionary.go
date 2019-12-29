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

// IsSetOrDictionary returns true if the string is a formatted set or dictionary.
func IsSetOrDictionary(s string) bool {
	return len(s) >= 2 && strings.HasPrefix(s, DictionaryPrefix) && strings.HasSuffix(s, DictionarySuffix)
}
