// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// FormatNodes formats an array of nodes to a string.
func FormatNodes(nodes []Node, quotes []string, pretty bool, tabs int) []string {
	values := make([]string, 0)
	for _, node := range nodes {
		values = append(values, node.Dfl(quotes, pretty, tabs))
	}
	return values
}
