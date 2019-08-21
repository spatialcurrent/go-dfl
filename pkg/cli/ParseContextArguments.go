// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cli

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

func ParseContextArguments(args []string, funcs dfl.FunctionMap, quotes []string) (map[string]interface{}, error) {
	ctx := map[string]interface{}{}
	for _, a := range args {
		if !strings.Contains(a, "=") {
			return ctx, errors.New("Context attribute \"" + a + "\" does not contain \"=\".")
		}
		pair := strings.SplitN(a, "=", 2)
		value, _, err := dfl.Parse(strings.TrimSpace(pair[1]))
		if err != nil {
			return ctx, errors.Wrap(err, "Could not parse context variable")
		}
		value = value.Compile()
		switch value.(type) {
		case dfl.Array:
			_, arr, err := value.(dfl.Array).Evaluate(map[string]interface{}{}, map[string]interface{}{}, funcs, quotes[1:])
			if err != nil {
				return ctx, errors.Wrap(err, "error evaluating context expression for "+strings.TrimSpace(pair[0]))
			}
			ctx[strings.TrimSpace(pair[0])] = arr
		case dfl.Set:
			_, arr, err := value.(dfl.Set).Evaluate(map[string]interface{}{}, map[string]interface{}{}, funcs, quotes[1:])
			if err != nil {
				return ctx, errors.Wrap(err, "error evaluating context expression for "+strings.TrimSpace(pair[0]))
			}
			ctx[strings.TrimSpace(pair[0])] = arr
		case dfl.Literal:
			ctx[strings.TrimSpace(pair[0])] = value.(dfl.Literal).Value
		case *dfl.Literal:
			ctx[strings.TrimSpace(pair[0])] = value.(*dfl.Literal).Value
		default:
			ctx[strings.TrimSpace(pair[0])] = dfl.TryConvertString(pair[1])
		}
	}
	return ctx, nil
}
