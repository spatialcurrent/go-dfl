// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cli

import (
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-stringify/pkg/stringify"

	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

func ParseInitialVariables(str string, funcs dfl.FunctionMap, quotes []string) (map[string]interface{}, error) {
	vars := map[string]interface{}{}
	if len(str) > 0 {
		_, result, err := dfl.ParseCompileEvaluateMap(
			str,
			dfl.NoVars,
			dfl.NoContext,
			funcs,
			quotes)
		if err != nil {
			return vars, errors.Wrap(err, "error parsing initial dfl vars as map")
		}
		stringifiedResult, err := stringify.StringifyMapKeys(result, stringify.NewDefaultStringer())
		if err != nil {
			return vars, errors.Wrap(err, "error stringifying result")
		}
		if m, ok := stringifiedResult.(map[string]interface{}); ok {
			vars = m
		}
	}
	return vars, nil
}
