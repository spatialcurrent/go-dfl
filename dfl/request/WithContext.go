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

func WithContext(requestContext context.Context, dflContext interface{}) context.Context {
  return context.WithValue(requestContext, contextKeyContext, dflContext)
}
