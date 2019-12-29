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

// ErrorEvaluate is an error returned when an error occurs during evaluation of a Node.
type ErrorEvaluate struct {
	Node   Node     // the name of the Function
	Quotes []string // the quotes to use
}

// Error returns the error as a string.
func (e ErrorEvaluate) Error() string {
	//return fmt.Sprintf("error evaluating %s (%q)",  Name(e.Node), e.Node.Dfl(e.Quotes, false, 0))
	return fmt.Sprintf("error evaluating %s (%#v)", Name(e.Node), e.Node.Map())
}
