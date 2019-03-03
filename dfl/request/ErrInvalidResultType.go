// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
	"fmt"
	"reflect"
)

type ErrInvalidResultType struct {
	Type reflect.Type
}

func (e *ErrInvalidResultType) Error() string {
	return fmt.Sprintf("invalid result type %s", e.Type.Name())
}
