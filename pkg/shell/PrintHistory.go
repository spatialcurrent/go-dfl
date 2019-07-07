// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package shell

import (
	"fmt"
	"strconv"
)

func PrintHistory(history []string) {
	fmt.Println("# History")
	for i, line := range history {
		fmt.Println(strconv.Itoa(i+1) + "\t" + line)
	}
}
