// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// ParseList parses a list of values.
func ParseListOrKeyValue(in string) (bool, []Node, map[Node]Node, error) {

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
			(c == ',' || c == ':')) {
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
			if i+1 == len(in) {
				value, _, err := Parse(in)
				if err != nil {
					return true, make([]Node, 0), map[Node]Node{}, errors.Wrap(err, "error parsing single value in set")
				}
				return true, []Node{value}, map[Node]Node{}, nil
				//return true, []Node{&Literal{Value: TryConvertString(in)}}, map[Node]Node{}, nil
			} else if in[i+1] == ',' {
				nodes, err := ParseList(in)
				return true, nodes, map[Node]Node{}, err
			} else if in[i+1] == ':' {
				nodes, err := ParseKeyValue(in)
				return false, make([]Node, 0), nodes, err
			}
		}
	}

	if leftparentheses > rightparentheses {
		return true, make([]Node, 0), map[Node]Node{}, errors.New("too few closing parentheses " + fmt.Sprint(leftparentheses) + " | " + fmt.Sprint(rightparentheses))
	} else if leftparentheses < rightparentheses {
		return true, make([]Node, 0), map[Node]Node{}, errors.New("too many closing parentheses " + fmt.Sprint(leftparentheses) + " | " + fmt.Sprint(rightparentheses))
	} else if leftcurlybrackets > rightcurlybrackets {
		return true, make([]Node, 0), map[Node]Node{}, errors.New("too few closing curly brackets " + fmt.Sprint(leftcurlybrackets) + " | " + fmt.Sprint(rightcurlybrackets))
	} else if leftcurlybrackets < rightcurlybrackets {
		return true, make([]Node, 0), map[Node]Node{}, errors.New("too many closing curly brackets " + fmt.Sprint(leftcurlybrackets) + " | " + fmt.Sprint(rightparentheses))
	} else if leftsquarebrackets > rightsquarebrackets {
		return true, make([]Node, 0), map[Node]Node{}, errors.New("too few closing square brackets " + fmt.Sprint(leftsquarebrackets) + " | " + fmt.Sprint(rightsquarebrackets))
	} else if leftsquarebrackets < rightsquarebrackets {
		return true, make([]Node, 0), map[Node]Node{}, errors.New("too many closing square brackets " + fmt.Sprint(leftsquarebrackets) + " | " + fmt.Sprint(rightsquarebrackets))
	}

	return true, make([]Node, 0), map[Node]Node{}, nil
}
