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

// Parse is the primary entrypoint for the DFL library.
// Parse takes a DFL expression string as input and returns an Abstract Synatax Tree (AST), and error if any.
func Parse(in string) (Node, error) {

	var root Node

	if len(in) == 0 {
		return root, errors.New("Error: Input string is empty.")
	}

	leftparentheses := 0
	rightparentheses := 0
	leftcurlybrackets := 0
	rightcurlybrackets := 0
	leftsquarebrackets := 0
	rightsquarebrackets := 0
	singlequotes := 0
	doublequotes := 0
	backticks := 0

	for i, c := range in {

		s := strings.TrimSpace(in[0 : i+1])
		s_lc := strings.ToLower(s)
		remainder := strings.TrimSpace(in[i+1:])

		if singlequotes == 0 && doublequotes == 0 && backticks == 0 {
			if c == '(' {
				leftparentheses += 1
			} else if c == ')' {
				rightparentheses += 1
			} else if c == '[' {
				leftsquarebrackets += 1
			} else if c == ']' {
				rightsquarebrackets += 1
			} else if c == '{' {
				leftcurlybrackets += 1
			} else if c == '}' {
				rightcurlybrackets += 1
			} else if c == '\'' {
				singlequotes += 1
			} else if c == '"' {
				doublequotes += 1
			} else if c == '`' {
				backticks += 1
			}
		} else if singlequotes == 1 && c == '\'' {
			singlequotes -= 1
		} else if doublequotes == 1 && c == '"' {
			doublequotes -= 1
		} else if backticks == 1 && c == '`' {
			backticks -= 1
		}

		if leftparentheses == rightparentheses &&
			leftsquarebrackets == rightsquarebrackets &&
			leftcurlybrackets == rightcurlybrackets &&
			singlequotes == 0 &&
			doublequotes == 0 &&
			backticks == 0 {
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

			} else if s_lc == "+" && in[i+1] != '=' {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Add{&BinaryOperator{Right: right}}, nil

			} else if s_lc == "-" && in[i+1] == ' ' { // space is require to exclude negative numbers and -=

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Subtract{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if s_lc == "/" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Divide{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, nil

			} else if s_lc == "|" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Pipe{&BinaryOperator{Right: right}}, nil

			} else if s_lc == ":=" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &Assign{&BinaryOperator{Right: right}}, nil

			} else if s_lc == "+=" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &AssignAdd{&BinaryOperator{Right: right}}, nil

			} else if s_lc == "-=" {

				right, err := Parse(remainder)
				if err != nil {
					return right, err
				}
				return &AssignSubtract{&BinaryOperator{Right: right}}, nil

			} else if len(remainder) == 0 || in[i+1] == ' ' || in[i+1] == '\n' {
				if IsQuoted(s) {
					return ParseLiteral(s[1:len(s)-1], remainder)
				} else if IsAttribute(s) {
					return ParseAttribute(s, remainder)
				} else if IsVariable(s) {
					return ParseVariable(s, remainder)
				} else if IsArray(s) {
					return ParseArray(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if IsSetOrDictionary(s) {
					return ParseSetOrDictionary(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if IsSub(s) {
					return ParseSub(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if IsFunction(s) {
					return ParseFunction(strings.TrimSpace(s), remainder)
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

				} else if s_lc == "within" || s_lc == "between" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &Within{&BinaryOperator{Right: right}}, nil

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
				} else {
					return ParseLiteral(TryConvertString(s), remainder)
				}
			} else {

			}
		}

	}

	return root, errors.New("Invalid expression syntax for \"" + in + "\".")
}
