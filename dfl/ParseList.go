// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"github.com/spatialcurrent/go-dfl/dfl/syntax"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// ParseList parses a list of values.
func ParseList(in string) ([]Node, error) {

	nodes := make([]Node, 0)

	singlequotes := 0
	doublequotes := 0
	backticks := 0

	leftparentheses := 0
	rightparentheses := 0
	leftcurlybrackets := 0
	rightcurlybrackets := 0
	leftsquarebrackets := 0
	rightsquarebrackets := 0

	in = strings.TrimSpace(in)
	s := ""

	for i, c := range in {

		// If you're note in a quoted string, then you can start one
		// If you're in a quoted string, only exit with matching quote.
		if singlequotes == 0 && doublequotes == 0 && backticks == 0 {
			switch c {
			case '\'':
				singlequotes += 1
			case '"':
				doublequotes += 1
			case '`':
				backticks += 1
			case '(':
				leftparentheses += 1
			case '[':
				leftsquarebrackets += 1
			case '{':
				leftcurlybrackets += 1
			case ')':
				rightparentheses += 1
			case ']':
				rightsquarebrackets += 1
			case '}':
				rightcurlybrackets += 1
			}
		} else if singlequotes == 1 && c == '\'' {
			singlequotes -= 1
		} else if doublequotes == 1 && c == '"' {
			doublequotes -= 1
		} else if backticks == 1 && c == '`' {
			backticks -= 1
		}

		// If not (within string/array/set/sub and c == ,)
		if !(singlequotes == 0 &&
			doublequotes == 0 &&
			backticks == 0 &&
			leftparentheses == rightparentheses &&
			leftcurlybrackets == rightcurlybrackets &&
			leftsquarebrackets == rightsquarebrackets &&
			c == ',') {
			s += string(c)
		}

		// If sub/array/set and string are closed.
		if singlequotes == 0 &&
			doublequotes == 0 &&
			backticks == 0 &&
			leftparentheses == rightparentheses &&
			leftsquarebrackets == rightsquarebrackets &&
			leftcurlybrackets == rightcurlybrackets {
			// If end of input or argument
			if i+1 == len(in) || in[i+1] == ',' {
				s = strings.TrimSpace(s)
				if syntax.IsQuoted(s) {
					nodes = append(nodes, &Literal{Value: s[1 : len(s)-1]})
				} else if syntax.IsAttribute(s) {
					attr, _, err := ParseAttribute(s, "")
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing attribute in list "+s)
					}
					nodes = append(nodes, attr)
				} else if syntax.IsVariable(s) {
					variable, _, err := ParseVariable(s, "")
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing variable in list "+s)
					}
					nodes = append(nodes, variable)
				} else if syntax.IsArray(s) {
					arr, _, err := ParseArray(strings.TrimSpace(s[1:len(s)-1]), "")
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing array in list "+s)
					}
					nodes = append(nodes, arr)
				} else if syntax.IsSetOrDictionary(s) {
					setOrDictionary, _, err := ParseSetOrDictionary(strings.TrimSpace(s[1:len(s)-1]), "")
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing set in list "+s)
					}
					nodes = append(nodes, setOrDictionary)
				} else if syntax.IsSub(s) {
					sub, _, err := ParseSub(strings.TrimSpace(s[1:len(s)-1]), "")
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing sub in list "+s)
					}
					nodes = append(nodes, sub)
				} else if syntax.IsFunction(s) {
					f, _, err := ParseFunction(s, "")
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing function in list "+s)
					}
					nodes = append(nodes, f)
				} else {
					nodes = append(nodes, &Literal{Value: TryConvertString(s)})
				}
				s = ""
			}
		}

	}

	if leftparentheses > rightparentheses {
		return nodes, errors.New("too few closing parentheses " + fmt.Sprint(leftparentheses) + " | " + fmt.Sprint(rightparentheses))
	} else if leftparentheses < rightparentheses {
		return nodes, errors.New("too many closing parentheses " + fmt.Sprint(leftparentheses) + " | " + fmt.Sprint(rightparentheses))
	} else if leftcurlybrackets > rightcurlybrackets {
		return nodes, errors.New("too few closing curly brackets " + fmt.Sprint(leftcurlybrackets) + " | " + fmt.Sprint(rightcurlybrackets))
	} else if leftcurlybrackets < rightcurlybrackets {
		return nodes, errors.New("too many closing curly brackets " + fmt.Sprint(leftcurlybrackets) + " | " + fmt.Sprint(rightparentheses))
	} else if leftsquarebrackets > rightsquarebrackets {
		return nodes, errors.New("too few closing square brackets " + fmt.Sprint(leftsquarebrackets) + " | " + fmt.Sprint(rightsquarebrackets))
	} else if leftsquarebrackets < rightsquarebrackets {
		return nodes, errors.New("too many closing square brackets " + fmt.Sprint(leftsquarebrackets) + " | " + fmt.Sprint(rightsquarebrackets))
	}

	return nodes, nil
}
