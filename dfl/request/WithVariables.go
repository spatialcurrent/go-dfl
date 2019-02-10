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

func WithVariables(requestContext context.Context, variables map[string]interface{}) context.Context {
	return context.WithValue(requestContext, contextKeyVariables, variables)
}
