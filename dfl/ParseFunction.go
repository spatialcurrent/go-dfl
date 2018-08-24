// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"strings"
	"unicode"
)

import (
	"github.com/pkg/errors"
)

// ParseFunction parses a function from in and attaches the remainder.
func ParseFunction(in string, remainder string) (Node, error) {

	functionName := ""
	arguments := make([]Node, 0)
	leftparentheses := 0
	rightparentheses := 0
	singlequotes := 0
	doublequotes := 0
	backticks := 0

	in = strings.TrimSpace(in)
	s := ""

	for i := 0; i < len(in); i++ {

		c := in[i]

		s += string(c)

		if c == '"' && singlequotes == 0 && backticks == 0 {
			doublequotes ^= 1
		} else if c == '\'' && doublequotes == 0 && backticks == 0 {
			singlequotes ^= 1
		} else if c == '`' && singlequotes == 0 && doublequotes == 0 {
			backticks ^= 1
		} else if c == '(' {
			leftparentheses += 1

			if leftparentheses == 1 {
				functionName = strings.TrimSpace(s[:len(s)-1])
				s = ""
			}

		} else if c == ')' {

			rightparentheses += 1

			if singlequotes == 0 && doublequotes == 0 && backticks == 0 {
				if leftparentheses-rightparentheses == 1 {
					// if leftparentheses-rightparentheses > 0 {
					subFunction, err := ParseFunction(strings.TrimSpace(s), "")
					if err != nil {
						return &Function{}, errors.Wrap(err, "error parsing sub function < "+strings.TrimSpace(s)+" >")
					}
					arguments = append(arguments, subFunction)
					s = ""
				} else if leftparentheses == rightparentheses {
					s = strings.TrimSpace(s[0 : len(s)-1])
					if len(s) > 0 {
						if len(s) >= 2 && ((strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) || (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")) || (strings.HasPrefix(s, "`") && strings.HasSuffix(s, "`"))) {
							arguments = append(arguments, &Literal{Value: s[1 : len(s)-1]})
						} else if strings.HasPrefix(strings.TrimLeftFunc(s, unicode.IsSpace), "@") {
							arguments = append(arguments, &Attribute{Name: strings.TrimLeftFunc(s, unicode.IsSpace)[1:]})
						} else if strings.Contains(s, "(") {
							subFunction, err := ParseFunction(strings.TrimSpace(s), "")
							if err != nil {
								return &Function{}, errors.Wrap(err, "error parsing subfunction as argument")
							}
							arguments = append(arguments, subFunction)
						} else {
							arguments = append(arguments, &Literal{Value: TryConvertString(s)})
						}
					}
					s = ""
				}
			}

			//s = ""

		} else if singlequotes == 0 && doublequotes == 0 && backticks == 0 && (leftparentheses-rightparentheses) == 1 && c == ',' {
			s = strings.TrimSpace(s[0 : len(s)-1])
			if len(s) > 0 {
				if len(s) >= 2 && ((strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) || (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")) || (strings.HasPrefix(s, "`") && strings.HasSuffix(s, "`"))) {
					arguments = append(arguments, &Literal{Value: s[1 : len(s)-1]})
				} else if strings.HasPrefix(strings.TrimLeftFunc(s, unicode.IsSpace), "@") {
					arguments = append(arguments, &Attribute{Name: strings.TrimLeftFunc(s, unicode.IsSpace)[1:]})
				} else {
					arguments = append(arguments, &Literal{Value: TryConvertString(s)})
				}
			}
			s = ""
		}

	}

	if leftparentheses > rightparentheses {
		return &Function{}, errors.New("too few closing parentheses " + fmt.Sprint(leftparentheses) + " | " + fmt.Sprint(rightparentheses))
	} else if leftparentheses < rightparentheses {
		return &Function{}, errors.New("too many closing parentheses " + fmt.Sprint(leftparentheses) + " | " + fmt.Sprint(rightparentheses))
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
