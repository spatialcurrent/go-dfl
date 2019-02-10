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

func WithExpressionAndContext(requestContext context.Context, expression string, dflContext interface{}) context.Context {
	return WithContext(WithExpression(requestContext, expression), dflContext)
}
