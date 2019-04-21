// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

// FormatNodes formats an array of nodes to a string.
func FormatNodes(nodes []Node, quotes []string, pretty bool, tabs int) []string {
	values := make([]string, 0)
	for _, node := range nodes {
		values = append(values, strings.TrimSpace(node.Dfl(quotes, pretty, tabs+1)))
	}
	return values
}
