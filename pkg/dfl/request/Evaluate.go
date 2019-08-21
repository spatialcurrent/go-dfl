// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
	"context"

	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-dfl/dfl"
	"github.com/spatialcurrent/go-dfl/pkg/dfl/cache"
)

func Evaluate(ctx context.Context) (map[string]interface{}, interface{}, error) {

	vars, ok := ctx.Value(contextKeyVariables).(map[string]interface{})
	if !ok {
		return map[string]interface{}{}, nil, &ErrMissingContextValue{Name: contextKeyVariables.String()}
	}

	funcs, ok := ctx.Value(contextKeyFunctions).(dfl.FunctionMap)
	if !ok {
		return vars, nil, &ErrMissingContextValue{Name: contextKeyFunctions.String()}
	}

	dflContext := ctx.Value(contextKeyContext)
	if ctx == nil {
		return vars, nil, &ErrMissingContextValue{Name: contextKeyContext.String()}
	}

	exp, ok := ctx.Value(contextKeyExpression).(string)
	if !ok {
		return vars, nil, &ErrMissingContextValue{Name: contextKeyExpression.String()}
	}

	if c, cacheExists := ctx.Value(contextKeyCache).(cache.Cache); cacheExists {
		n, err := c.ParseCompile(exp)
		if err != nil {
			return vars, nil, errors.Wrap(err, "error parsing and compiling ( "+exp+" )")
		}
		return n.Evaluate(vars, dflContext, funcs, dfl.DefaultQuotes)
	}

	n, err := dfl.ParseCompile(exp)
	if err != nil {
		return vars, nil, errors.Wrap(err, "error parsing and compiling ( "+exp+" )")
	}
	return n.Evaluate(vars, dflContext, funcs, dfl.DefaultQuotes)
}
