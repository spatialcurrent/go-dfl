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

func FormatList(values []string, delim string, pretty bool, tabs int) string {
	if pretty {
		return strings.Repeat(DefaultTab, tabs) + strings.Join(values, delim+"\n"+strings.Repeat(DefaultTab, tabs))
	}
	return strings.Join(values, delim+" ")
}
