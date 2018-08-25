// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
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
func ParseFunction(in string, remainder string) (Node, error) {

	s := strings.TrimSpace(in)
	for i, c := range s {

		if c == '(' {
			functionName := s[:i]
			arguments, err := ParseList(s[i+1 : len(s)-1])
			if err != nil {
				return &Function{}, errors.Wrap(err, "error parsing function arguments from "+s[i+1:len(s)-2])
			}

			if len(remainder) == 0 {
				return &Function{Name: functionName, Arguments: arguments}, nil
			}

			root, err := Parse(remainder)
			if err != nil {
				return root, err
			}

			err = AttachLeft(root, &Function{Name: functionName, Arguments: arguments})
			if err != nil {
				return root, errors.Wrap(err, "error attaching left for "+in)
			}

			return root, nil
		}
	}

	return &Function{}, errors.New("no left parentheses found when parsing function string " + in)
}
