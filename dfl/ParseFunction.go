// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

import (
	"github.com/pkg/errors"
)

// ParseFunction parses a function from in and attaches the remainder.
func ParseFunction(in string, remainder string) (Node, string, error) {

	s := strings.TrimSpace(in)
	for i, c := range s {

		if c == '(' {
			functionName := s[:i]
			arguments, err := ParseList(s[i+1 : len(s)-1])
			if err != nil {
				return &Function{}, remainder, errors.Wrap(err, "error parsing function arguments from "+s[i+1:len(s)-2])
			}

			if len(remainder) == 0 || (len(remainder) >= 2 && remainder[0] == ':' && remainder[1] != '=') {
				return &Function{Name: functionName, MultiOperator: &MultiOperator{Arguments: arguments}}, remainder, nil
			}

			root, remainder, err := Parse(remainder)
			if err != nil {
				return root, remainder, err
			}

			err = AttachLeft(root, &Function{Name: functionName, MultiOperator: &MultiOperator{Arguments: arguments}})
			if err != nil {
				return root, remainder, errors.Wrap(err, "error attaching left for "+in)
			}

			return root, remainder, nil
		}
	}

	return &Function{}, remainder, errors.New("no left parentheses found when parsing function string " + in)
}
