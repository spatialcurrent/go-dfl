// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
  "context"
)

import (
  "github.com/spatialcurrent/go-dfl/dfl"
)

func WithFunctions(requestContext context.Context, funcs dfl.FunctionMap) context.Context {
  return context.WithValue(requestContext, contextKeyFunctions, funcs)
}
