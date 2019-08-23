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

func ClearLine() {
	fmt.Print("\r")     // carriage return to beginning of line
	fmt.Print("\033[K") // https://www.student.cs.uwaterloo.ca/~cs452/terminal.html
}
