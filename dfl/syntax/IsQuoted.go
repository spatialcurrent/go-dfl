// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package syntax

// IsQuoted returns true if the string is a quoted string (using single quotes, double quotes, or backticks)
func IsQuoted(s string) bool {
	if len(s) < 2 {
		return false
	}
	return (s[0] == '\'' && s[len(s)-1] == '\'') || (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '`' && s[len(s)-1] == '`')
}
