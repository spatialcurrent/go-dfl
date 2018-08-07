// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
	"unicode"
)

import (
	"github.com/pkg/errors"
)

// ParseSet parses a Set Node and recursively any remainder.
// If parameter "in" is gramatically a child node, then return the parent node.
// DFL sets can include Attribute or Literal Nodes.
// As all attribute references must start with an "@" character, parantheses are optional for literals except if a comma exists.
// Below are some example inputs
//
//	{bank, bureau_de_change, atm}
//	{1, 2, @target}
//	{Taco, Tacos, Burrito, Burritos, "Mexican Food", @example}
func ParseSet(in string, remainder string) (Node, error) {

	nodes := make([]Node, 0)
	singlequotes := 0
	doublequotes := 0

	in = strings.TrimSpace(in)
	s := ""

	for i, c := range in {

		if !(singlequotes == 0 && doublequotes == 0 && c == ',') {
			s += string(c)
			if c == '\'' && doublequotes == 0 {
				if singlequotes == 0 {
					singlequotes += 1
				} else {
					singlequotes -= 1
				}
			} else if c == '"' && singlequotes == 0 {
				if doublequotes == 0 {
					doublequotes += 1
				} else {
					doublequotes -= 1
				}
			}
		}

		if singlequotes == 0 && doublequotes == 0 && (i+1 == len(in) || in[i+1] == ',') {
			s = strings.TrimSpace(s)
			if len(s) >= 2 && ((strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) || (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\""))) {
				nodes = append(nodes, &Literal{Value: s[1 : len(s)-1]})
			} else if strings.HasPrefix(strings.TrimLeftFunc(s, unicode.IsSpace), "@") {
				nodes = append(nodes, &Attribute{Name: strings.TrimLeftFunc(s, unicode.IsSpace)[1:]})
			} else {
				nodes = append(nodes, &Literal{Value: TryConvertString(s)})
			}
			s = ""
		}

	}

	if len(remainder) == 0 {
		return &Set{Nodes: nodes}, nil
	}

	left := &Set{Nodes: nodes}
	root, err := Parse(remainder)
	if err != nil {
		return root, err
	}

	err = AttachLeft(root, left)
	if err != nil {
		return root, errors.Wrap(err, "error attaching left for set "+in)
	}

	return root, nil

}
