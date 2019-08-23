// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
	"context"

	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

func WithFunctions(requestContext context.Context, funcs dfl.FunctionMap) context.Context {
	return context.WithValue(requestContext, contextKeyFunctions, funcs)
}
