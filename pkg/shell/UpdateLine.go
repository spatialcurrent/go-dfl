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

import (
	"github.com/spf13/viper"
)

func UpdateLine(v *viper.Viper, vars map[string]interface{}, quotes []string, line string) error {

	varsBytes, err := FormatOutput(v, vars, quotes)
	if err != nil {
		return err
	}

	fmt.Print("(" + string(varsBytes) + ") > " + line)

	return nil
}
