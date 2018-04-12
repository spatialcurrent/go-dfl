package dfl

import (
	"regexp"
	"strings"
	"unicode"
)

import (
	"github.com/pkg/errors"
)

func ParseValue(s string, remainder string) (Node, error) {

	if len(remainder) == 0 {
		return &Literal{Value: s}, nil
	}

	left := &Literal{Value: s}
	root, err := Parse(remainder)
	if err != nil {
		return root, err
	}
	switch root.(type) {
	case *And:
		root.(*And).Left = left
	case *Or:
		root.(*Or).Left = left
	case *In:
		root.(*In).Left = left
	case *Like:
		root.(*Like).Left = left
	case *ILike:
		root.(*ILike).Left = left
	case *LessThan:
		root.(*LessThan).Left = left
	case *LessThanOrEqual:
		root.(*LessThanOrEqual).Left = left
	case *GreaterThan:
		root.(*GreaterThan).Left = left
	case *GreaterThanOrEqual:
		root.(*GreaterThanOrEqual).Left = left
	default:
		return root, errors.New("Invalid expression syntax for "+s+".  Root is not a binary operator")
	}
	return root, nil
}

func ParseAttribute(in string) (Node, error) {

	end := strings.Index(strings.TrimLeftFunc(in, unicode.IsSpace), " ")
	if end == -1 {
		return &Attribute{Name:strings.TrimSpace(in)[1:]}, nil
	}

	if len(strings.TrimSpace(in[end:])) == 0 {
		return &Attribute{Name:in[1:end]}, nil
	}

	left := &Attribute{Name: in[1:end]}
	root, err := Parse(in[end:])
	if err != nil {
		return root, err
	}
	switch root.(type) {
	case *And:
		root.(*And).Left = left
	case *Or:
		root.(*Or).Left = left
	case *In:
		root.(*In).Left = left
	case *Like:
		root.(*Like).Left = left
	case *ILike:
		root.(*ILike).Left = left
	case *LessThan:
		root.(*LessThan).Left = left
	case *LessThanOrEqual:
		root.(*LessThanOrEqual).Left = left
	case *GreaterThan:
		root.(*GreaterThan).Left = left
	case *GreaterThanOrEqual:
		root.(*GreaterThanOrEqual).Left = left
	default:
		return root, errors.New("Invalid expression syntax for "+in+".  Root is not a binary operator")
	}
	return root, nil

}

func ParseSub(s string, remainder string) (Node, error) {

	if len(remainder) == 0 {
		return Parse(s)
	}

  var root Node
	left, err := Parse(s)
	if err != nil {
		return root, err
	}

	root, err = Parse(remainder)
	if err != nil {
		return root, err
	}

	switch root.(type) {
	case *And:
		root.(*And).Left = left
	case *Or:
		root.(*Or).Left = left
	case *In:
		root.(*In).Left = left
	case *Like:
		root.(*Like).Left = left
	case *ILike:
		root.(*ILike).Left = left
	case *LessThan:
		root.(*LessThan).Left = left
	case *LessThanOrEqual:
		root.(*LessThanOrEqual).Left = left
	case *GreaterThan:
		root.(*GreaterThan).Left = left
	case *GreaterThanOrEqual:
		root.(*GreaterThanOrEqual).Left = left
	default:
		return root, errors.New("Invalid expression syntax for "+s+".  Root is not a binary operator")
	}

	return root, nil
}

func Parse(in string) (Node, error) {
  var root Node

	if len(in) == 0 {
		return root, errors.New("Error: Input string is empty.")
	}

  re, err := regexp.Compile("(\\s*)(?P<name>([a-zA-Z_\\d]+))(\\s*)\\((\\s*)(?P<args>(.)*?)(\\s*)\\)(\\s*)")
  if err != nil {
    return root, err
  }

	if strings.HasPrefix(strings.TrimLeftFunc(in, unicode.IsSpace), "@") {
		return ParseAttribute(in)
	} else {

		parentheses := 0
		for i, c := range in {

	    s := strings.TrimSpace(in[0:i+1])
	    s_lc := strings.ToLower(s)
	    remainder := strings.TrimSpace(in[i+1:])

			if c == '(' {
				parentheses += 1
			} else if c == ')' {
				parentheses -= 1
			}

      if parentheses == 0 && (len(remainder) == 0 || in[i+1] == ' ') {
				if len(s) >= 2  && ((strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'")) || (strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\""))) {
					return ParseValue(s[1:len(s) - 1], remainder)
				} else if len(s) >= 2  && strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
					return ParseSub(s[1: len(s) - 1], remainder)
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

				} else if s_lc == "<" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &LessThan{&BinaryOperator{Right: right}}, nil

				} else if s_lc == "<=" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &LessThanOrEqual{&BinaryOperator{Right: right}}, nil

				} else if s_lc == ">" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &GreaterThan{&BinaryOperator{Right: right}}, nil

				} else if s_lc == ">=" {

					right, err := Parse(remainder)
					if err != nil {
						return right, err
					}
					return &GreaterThanOrEqual{&BinaryOperator{Right: right}}, nil

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

				} else if re.MatchString(s) {
					return ParseFunction(s, remainder, re)
				}
			}

	  }
	}

  return root, errors.New("Invalid expression syntax for \""+in+"\".")
}
