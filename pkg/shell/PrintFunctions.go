// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package shell

import (
	"fmt"
	"sort"
	"strings"
)

import (
	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

func PrintFunctions(funcs dfl.FunctionMap) {
	fmt.Println("# Functions")
	lines := make([]string, 0, len(funcs))
	for k, _ := range funcs {
		lines = append(lines, k)
	}
	sort.Strings(lines)
	fmt.Println(strings.Join(lines, "\n"))
}
