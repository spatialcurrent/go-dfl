// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

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
			}
		}
	}
	return in
}
