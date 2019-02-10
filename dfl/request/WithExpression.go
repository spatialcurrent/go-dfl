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

func WithExpression(ctx context.Context, expression string) context.Context {
  return context.WithValue(ctx, contextKeyExpression, expression)
}
