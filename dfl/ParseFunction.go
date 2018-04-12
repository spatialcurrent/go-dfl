package dfl

import (
	"regexp"
	"strings"
)

import (
	"github.com/pkg/errors"
)

func ParseFunctionArguments(in string) ([]string, error) {

	args := []string{}

	re2, err := regexp.Compile("(\\s*)(?P<value>((\"([^\"]+?)\")|([^,\\s]+)))(\\s*)")
	if err != nil {
		return args, err
	}

	for _, m2 := range re2.FindAllStringSubmatch(in, -1) {
		g2 := map[string]string{}
		for i, name := range re2.SubexpNames() {
			if i != 0 {
				g2[name] = m2[i]
			}
		}
		if value, ok := g2["value"]; ok {
			value = strings.TrimSpace(value)
			if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
				args = append(args, value[1:len(value)-1])
			} else {
				args = append(args, value)
			}
		}
	}

	return args, nil
}

func ParseFunction(s string, remainder string, re *regexp.Regexp) (Node, error) {
	var root Node

	m := re.FindStringSubmatch(s)
	g := map[string]string{}
	for j, name := range re.SubexpNames() {
		if j != 0 {
			g[name] = m[j]
		}
	}
	name := g["name"]

	args, err := ParseFunctionArguments(g["args"])
	if err != nil {
		return root, err
	}

	if len(remainder) == 0 {
		return &Function{Name: name, Arguments: args}, nil
	}

	left := &Function{Name: name, Arguments: args}
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
		return root, errors.New("Invalid expression syntax for " + s + ".  Root is not a binary operator")
	}

	return root, nil
}
