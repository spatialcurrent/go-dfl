// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package shell

import (
	"encoding/json"
	"fmt"
)

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

import (
	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

func FormatOutput(v *viper.Viper, object interface{}, quotes []string) ([]byte, error) {
	if v.GetBool("json") {
		if v.GetBool("pretty") {
			outputBytes, err := json.MarshalIndent(object, "", "  ")
			if err != nil {
				return make([]byte, 0), errors.Wrap(err, "error marshalling result")
			}
			return outputBytes, nil
		}
		outputBytes, err := json.Marshal(object)
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error marshalling result")
		}
		return outputBytes, nil
	}

	if v.GetBool("yaml") {
		outputBytes, err := yaml.Marshal(object)
		if err != nil {
			return make([]byte, 0), errors.Wrap(err, "error marshalling result")
		}
		return outputBytes, nil
	}

	if v.GetBool("go") {
		return []byte(fmt.Sprintf("%#v", object)), nil
	}

	return []byte(dfl.TryFormatLiteral(object, quotes, v.GetBool("pretty"), v.GetInt("tabs"))), nil
}
