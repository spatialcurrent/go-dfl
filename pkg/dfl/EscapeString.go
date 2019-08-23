// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

func EscapeString(in string) string {
	for i, c := range in {
		switch c {
		case '\\':
			return in[0:i] + "\\\\" + EscapeString(in[i+1:])
		case '\n':
			return in[0:i] + "\\n" + EscapeString(in[i+1:])
		case '\t':
			return in[0:i] + "\\t" + EscapeString(in[i+1:])
		}
	}
	return in
}
