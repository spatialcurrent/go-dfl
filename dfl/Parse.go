// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	//"fmt"
	"regexp"
	"strings"
	//"unicode"
)

import (
	"github.com/pkg/errors"
)

// Parse is the primary entrypoint for the DFL library.
// Parse takes a DFL expression string as input and returns an Abstract Synatax Tree (AST), and error if any.
func Parse(in string) (Node, error) {

	var root Node

	if len(in) == 0 {
		return root, errors.New("Error: Input string is empty.")
	}

	re, err := regexp.Compile("(\\s*)(?P<name>([a-zA-Z_\\d]+))(\\s*)\\((\\s*)(?P<args>(.)*?)(\\s*)\\)(\\s*)")
	if err != nil {
		return root, err
	}

	parentheses := 0
	curlybrackets := 0
	squarebrackets := 0
	singlequotes := 0
	doublequotes := 0
	for i, c := range in {

		s := strings.TrimSpace(in[0 : i+1])
		s_lc := strings.ToLower(s)
		remainder := strings.TrimSpace(in[i+1:])

		if c == '(' {
			parentheses += 1
		} else if c == ')' {
			parentheses -= 1
		} else if singlequotes == 0 && doublequotes == 0 {
			if squarebrackets == 0 && c == '[' {
				squarebrackets += 1
			} else if squarebrackets == 1 && c == ']' {
				squarebrackets -= 1
			} else if curlybrackets == 0 && c == '{' {
				curlybrackets += 1
			} else if curlybrackets == 1 && c == '}' {
				curlybrackets -= 1
			} else if c == '\'' {
				singlequotes += 1
			} else if c == '"' {
				doublequotes += 1
			}
		} else if singlequotes == 1 && c == '\'' {
			singlequotes -= 1
		} else if doublequotes == 1 && c == '"' {
			doublequotes -= 1
		}

		if parentheses == 0 &&
			squarebrackets == 0 &&
			curlybrackets == 0 &&
			singlequotes == 0 &&
			doublequotes == 0 {
			if s_lc == "?:" {
				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Coalesce{&BinaryOperator{Right: right}}, nil
			} else if s_lc == ">=" {
				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &GreaterThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil
			} else if s_lc == "<" && in[i+1] != '=' {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &LessThan{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if s_lc == "<=" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &LessThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if s_lc == ">" && in[i+1] != '=' {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &GreaterThan{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if s_lc == ">=" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &GreaterThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if s_lc == "==" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Equal{&BinaryOperator{Right: right}}, nil

			} else if s_lc == "!=" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &NotEqual{&BinaryOperator{Right: right}}, nil

			} else if s_lc == "+" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Add{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if s_lc == "-" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Subtract{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if len(remainder) == 0 || in[i+1] == ' ' || in[i+1] == '\n' {
				if len(s) >= 2 && ((strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) || (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\""))) {
					return ParseLiteral(s[1:len(s)-1], remainder)
				} else if strings.HasPrefix(s, "@") {
					return ParseAttribute(s, remainder)
				} else if len(s) >= 2 && strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
					return ParseArray(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if len(s) >= 2 && strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
					return ParseSet(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if len(s) >= 2 && strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
					return ParseSub(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if s_lc == "and" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &And{&BinaryOperator{Right: right}}, nil

				} else if s_lc == "or" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &Or{&BinaryOperator{Right: right}}, nil

				} else if s_lc == "xor" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &Xor{&BinaryOperator{Right: right}}, nil

				} else if s_lc == "not" {

					node, err := Parse(remainder)
					if err != nil {
						return node, err
					}
					return &Not{&UnaryOperator{Node: node}}, nil

				} else if s_lc == "in" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &In{&BinaryOperator{Right: right}}, nil

				} else if s_lc == "like" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &Like{&BinaryOperator{Right: right}}, nil

				} else if s_lc == "ilike" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &ILike{&BinaryOperator{Right: right}}, nil

				} else if s_lc == "before" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &Before{&TemporalBinaryOperator{&BinaryOperator{Right: right}}}, nil

				} else if s_lc == "after" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &After{&TemporalBinaryOperator{&BinaryOperator{Right: right}}}, nil

				} else if re.MatchString(s) {
					return ParseFunction(s, remainder)
				} else {
					return ParseLiteral(TryConvertString(s), remainder)
				}
			} else {

			}
		}

	}

	return root, errors.New("Invalid expression syntax for \"" + in + "\".")
}
