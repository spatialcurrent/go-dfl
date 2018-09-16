// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// FormatNodes formats an array of nodes to a string.
func FormatNodes(nodes []Node, delim string, quotes []string, pretty bool, tabs int) string {
	str := ""
	for i, arg := range nodes {
		if i > 0 {
			str += delim
		}
		str += arg.Dfl(quotes, pretty, tabs)
	}
	return str
}
