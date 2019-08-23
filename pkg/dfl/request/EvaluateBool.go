// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
	"context"
	"reflect"
)

func EvaluateBool(ctx context.Context) (map[string]interface{}, bool, error) {

	vars, result, err := Evaluate(ctx)
	if err != nil {
		return vars, false, err
	}

	value, ok := result.(bool)
	if !ok {
		return vars, false, &ErrInvalidResultType{Type: reflect.TypeOf(result)}
	}

	return vars, value, nil
}
