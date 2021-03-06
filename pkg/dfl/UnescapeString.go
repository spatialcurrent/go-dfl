// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strconv"
)

// UnescapeString unescapes a string
//
//	- \\ => \
//	- \n => new line
//	- \r => carriage return
//	- \t => horizontal tab
//	- \s => space
//	- \u1234 => unicode value
func UnescapeString(in string) string {
	for i, c := range in {
		if c == '\\' {
			switch in[i+1] {
			case '\\':
				return in[0:i] + "\\" + UnescapeString(in[i+2:])
			case 'n':
				return in[0:i] + "\n" + UnescapeString(in[i+2:])
			case 'r':
				return in[0:i] + "\r" + UnescapeString(in[i+2:])
			case 't':
				return in[0:i] + "\t" + UnescapeString(in[i+2:])
			case 's':
				return in[0:i] + " " + UnescapeString(in[i+2:])
			case 'u':
				v, _, _, err := strconv.UnquoteChar("\\u"+in[i+2:i+6], '"')
				if err == nil {
					return in[0:i] + string(v) + UnescapeString(in[i+6:])
				}
			}
		}
	}
	return in
}
