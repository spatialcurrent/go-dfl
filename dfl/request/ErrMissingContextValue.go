// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package request

import (
  "fmt"
)

type ErrMissingContextValue struct {
  Name string
}

func (e *ErrMissingContextValue) Error() string {
  return fmt.Sprintf("missing context value for name %s", e.Name)
}
