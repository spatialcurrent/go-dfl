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

	"github.com/spatialcurrent/go-dfl/pkg/dfl/syntax"
)

func parseKeyOrValueString(s string) (Node, error) {
	if syntax.IsQuoted(s) {
		return &Literal{Value: UnescapeString(s[1 : len(s)-1])}, nil
	} else if syntax.IsAttribute(s) {
		attr, _, err := ParseAttribute(s, "")
		if err != nil {
			return attr, errors.Wrap(err, "error parsing attribute in list "+s)
		}
		return attr, nil
	} else if syntax.IsVariable(s) {
		variable, _, err := ParseVariable(s, "")
		if err != nil {
			return variable, errors.Wrap(err, "error parsing variable in list "+s)
		}
		return variable, nil
	} else if syntax.IsArray(s) {
		arr, _, err := ParseArray(strings.TrimSpace(s[1:len(s)-1]), "")
		if err != nil {
			return arr, errors.Wrap(err, "error parsing array in list "+s)
		}
		return arr, nil
	} else if syntax.IsSetOrDictionary(s) {
		setOrDictionary, _, err := ParseSetOrDictionary(strings.TrimSpace(s[1:len(s)-1]), "")
		if err != nil {
			return setOrDictionary, errors.Wrap(err, "error parsing set in list "+s)
		}
		return setOrDictionary, nil
	} else if syntax.IsSub(s) {
		sub, _, err := ParseSub(strings.TrimSpace(s[1:len(s)-1]), "")
		if err != nil {
			return sub, errors.Wrap(err, "error parsing sub in list "+s)
		}
		return sub, nil
	} else if syntax.IsFunction(s) {
		f, _, err := ParseFunction(s, "")
		if err != nil {
			return f, errors.Wrap(err, "error parsing function in list "+s)
		}
		return f, nil
	}

	return &Literal{Value: TryConvertString(s)}, nil
}

// ParseKeyValue parses a sequence of key value pairs
func ParseKeyValue(in string) (map[Node]Node, error) {

	nodes := map[Node]Node{}

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
	key := ""

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
			if i+1 == len(in) || in[i+1] == ',' || in[i+1] == ':' {
				if i+1 != len(in) && in[i+1] == ':' {
					key = strings.TrimSpace(s)
				} else if len(key) > 0 {
					keyNode, err := parseKeyOrValueString(key)
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing key for key-pair")
					}
					valueNode, err := parseKeyOrValueString(strings.TrimSpace(s))
					if err != nil {
						return nodes, errors.Wrap(err, "error parsing value for key-pair")
					}
					nodes[keyNode] = valueNode
					key = ""
				} else {
					return nodes, errors.New("missing key when parsing key-value pairs")
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
