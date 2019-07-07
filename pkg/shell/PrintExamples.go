// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package shell

import (
	"fmt"
)

func PrintExamples() {
	fmt.Println("# Examples")
	fmt.Println(`
# Add Numbers
$x := 10
$y := 20
$x + $y
`)
}
