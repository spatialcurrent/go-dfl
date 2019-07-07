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

func PrintHelp() {
	fmt.Println("# Help")
	fmt.Println("# examples: print examples")
	fmt.Println("# history: print history")
	fmt.Println("# funcs, functions: print function names")
	fmt.Println("# Ctrl-C, exit, quit: exit shell")
	fmt.Println("# Ctrl-X: cut current line")
	fmt.Println("# Ctrl-V: paste from clipboard")
}
