// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
)

// ErrParseTernary is an error when parsing ternary expressions.
type ErrParseTernary struct {
	Original  string // original
	Condition string // the condition
	True      string // the true value
	False     string // the false value
}

// Error returns the error as a string.
func (e ErrParseTernary) Error() string {
	return fmt.Sprintf("error parsing ternary expression: original = %q | condition = %q | true = %q | false = %q", e.Original, e.Condition, e.True, e.False)
}
