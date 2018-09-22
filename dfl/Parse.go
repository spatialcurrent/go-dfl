// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/spatialcurrent/go-dfl/dfl/syntax"
	"strings"
)

import (
	"github.com/pkg/errors"
)

// Parse parses a DFL expression into an an Abstract Synatax Tree (AST).
// Parse returns the AST, remainder, and error if any.
func Parse(in string) (Node, string, error) {

	var root Node

	if len(in) == 0 {
		return root, "", errors.New("Error: Input string is empty.")
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
			if s_lc == "?" && in[i+1] != '.' && in[i+1] != ':' {

				t, remainder, err := Parse(strings.TrimSpace(remainder))
				if err != nil {
					return t, remainder, err
				}

				f, remainder, err := Parse(strings.TrimSpace(remainder[1:]))
				if err != nil {
					return f, remainder, err
				}

				return &TernaryOperator{True: t, False: f}, remainder, nil

			} else if s_lc == "?:" {
				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Coalesce{&BinaryOperator{Right: right}}, remainder, nil
			} else if s_lc == ">=" {
				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &GreaterThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil
			} else if s_lc == "<" && in[i+1] != '=' {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &LessThan{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

			} else if s_lc == "<=" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &LessThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

			} else if s_lc == ">" && in[i+1] != '=' {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &GreaterThan{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

			} else if s_lc == ">=" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &GreaterThanOrEqual{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

			} else if s_lc == "==" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Equal{&BinaryOperator{Right: right}}, remainder, nil

			} else if s_lc == "!=" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &NotEqual{&BinaryOperator{Right: right}}, remainder, nil

			} else if s_lc == "+" && in[i+1] != '=' {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Add{&BinaryOperator{Right: right}}, remainder, nil

			} else if s_lc == "-" && in[i+1] == ' ' { // space is require to exclude negative numbers and -=

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Subtract{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

			} else if s_lc == "/" && in[i+1] != '=' {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Divide{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

			} else if s_lc == "*" && in[i+1] != '=' {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Multiply{&NumericBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

			} else if s_lc == "|" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Pipe{&BinaryOperator{Right: right}}, remainder, nil

			} else if s_lc == ":=" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &Assign{&BinaryOperator{Right: right}}, remainder, nil

			} else if s_lc == "+=" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &AssignAdd{&BinaryOperator{Right: right}}, remainder, nil

			} else if s_lc == "-=" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &AssignSubtract{&BinaryOperator{Right: right}}, remainder, nil

			} else if s_lc == "*=" {

				right, remainder, err := Parse(remainder)
				if err != nil {
					return right, remainder, err
				}
				return &AssignMultiply{&BinaryOperator{Right: right}}, remainder, nil

			} else if len(remainder) == 0 || in[i+1] == ' ' || in[i+1] == '\n' {
				if syntax.IsQuoted(s) {
					return ParseLiteral(s[1:len(s)-1], remainder)
				} else if syntax.IsAttribute(s) {
					return ParseAttribute(s, remainder)
				} else if syntax.IsVariable(s) {
					return ParseVariable(s, remainder)
				} else if syntax.IsArray(s) {
					return ParseArray(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if syntax.IsSetOrDictionary(s) {
					return ParseSetOrDictionary(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if syntax.IsSub(s) {
					return ParseSub(strings.TrimSpace(s[1:len(s)-1]), remainder)
				} else if syntax.IsFunction(s) {
					return ParseFunction(strings.TrimSpace(s), remainder)
				} else if s_lc == "and" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &And{&BinaryOperator{Right: right}}, remainder, nil

				} else if s_lc == "or" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &Or{&BinaryOperator{Right: right}}, remainder, nil

				} else if s_lc == "xor" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &Xor{&BinaryOperator{Right: right}}, remainder, nil

				} else if s_lc == "not" || (s_lc == "!" && in[i+1] != '=') {

					node, remainder, err := Parse(remainder)
					if err != nil {
						return node, remainder, err
					}
					return &Not{&UnaryOperator{Node: node}}, remainder, nil

				} else if s_lc == "in" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &In{&BinaryOperator{Right: right}}, remainder, nil

				} else if s_lc == "within" || s_lc == "between" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &Within{&BinaryOperator{Right: right}}, remainder, nil

				} else if s_lc == "like" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &Like{&BinaryOperator{Right: right}}, remainder, nil

				} else if s_lc == "ilike" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &ILike{&BinaryOperator{Right: right}}, remainder, nil

				} else if s_lc == "before" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &Before{&TemporalBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil

				} else if s_lc == "after" {

					right, remainder, err := Parse(remainder)
					if err != nil {
						return right, remainder, err
					}
					return &After{&TemporalBinaryOperator{&BinaryOperator{Right: right}}}, remainder, nil
				} else {
					return ParseLiteral(TryConvertString(s), strings.TrimSpace(remainder))
				}
			} else {

			}
		}

	}

	return root, "", errors.New("Invalid expression syntax for \"" + in + "\".")
}
