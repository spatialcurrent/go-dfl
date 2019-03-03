// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
	"context"
)

func GetExpression(ctx context.Context) (string, error) {
	exp, ok := ctx.Value(contextKeyExpression).(string)
	if !ok {
		return "", &ErrMissingContextValue{Name: contextKeyExpression.String()}
	}
	return exp, nil
}
