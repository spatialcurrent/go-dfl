// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package builder

import (
	"strings"
	"unicode"
)

var indentSpace = "  "

type Node interface {
	Dfl(quotes []string, pretty bool, tabs int) string
	Sql(pretty bool, tabs int) string
}

type Builder struct {
	indent    int
	quotes    []string
	pretty    bool
	tabs      int
	left      Node
	operator  string
	right     Node
	trimRight bool
}

func New(quotes []string, tabs int) Builder {
	return Builder{
		indent:    0,
		quotes:    quotes,
		pretty:    false,
		tabs:      tabs,
		left:      nil,
		operator:  "",
		right:     nil,
		trimRight: false,
	}
}

func (b Builder) Pretty(pretty bool) Builder {
	return Builder{
		indent:    b.indent,
		quotes:    b.quotes,
		pretty:    b.pretty,
		tabs:      b.tabs,
		left:      b.left,
		operator:  b.operator,
		right:     b.right,
		trimRight: b.trimRight,
	}
}

func (b Builder) Indent(indent int) Builder {
	return Builder{
		indent:    indent,
		quotes:    b.quotes,
		pretty:    b.pretty,
		tabs:      b.tabs,
		left:      b.left,
		operator:  b.operator,
		right:     b.right,
		trimRight: b.trimRight,
	}
}

func (b Builder) Tabs(tabs int) Builder {
	return Builder{
		indent:    b.indent,
		quotes:    b.quotes,
		pretty:    b.pretty,
		tabs:      tabs,
		left:      b.left,
		operator:  b.operator,
		right:     b.right,
		trimRight: b.trimRight,
	}
}

func (b Builder) Left(n Node) Builder {
	return Builder{
		indent:    b.indent,
		quotes:    b.quotes,
		pretty:    b.pretty,
		tabs:      b.tabs,
		left:      n,
		operator:  b.operator,
		right:     b.right,
		trimRight: b.trimRight,
	}
}

func (b Builder) Op(operator string) Builder {
	return Builder{
		indent:    b.indent,
		quotes:    b.quotes,
		pretty:    b.pretty,
		tabs:      b.tabs,
		left:      b.left,
		operator:  operator,
		right:     b.right,
		trimRight: b.trimRight,
	}
}

func (b Builder) Right(n Node) Builder {
	return Builder{
		indent:    b.indent,
		quotes:    b.quotes,
		pretty:    b.pretty,
		tabs:      b.tabs,
		left:      b.left,
		operator:  b.operator,
		right:     n,
		trimRight: b.trimRight,
	}
}

func (b Builder) TrimRight(trimRight bool) Builder {
	return Builder{
		indent:    b.indent,
		quotes:    b.quotes,
		pretty:    b.pretty,
		tabs:      b.tabs,
		left:      b.left,
		operator:  b.operator,
		right:     b.right,
		trimRight: trimRight,
	}
}

func (b Builder) Dfl() string {
	str := ""
	if b.indent > 0 {
		str += strings.Repeat(indentSpace, b.indent)
	}
	str += "("
	if b.pretty {
		str += "\n"
	}
	if b.left != nil {
		str += b.left.Dfl(b.quotes, b.pretty, b.tabs)
		str += " "
	}
	str += b.operator
	if b.pretty && b.left != nil {
		str += "\n"
	} else {
		str += " "
	}
	if b.trimRight {
		str += strings.TrimLeftFunc(b.right.Dfl(b.quotes, b.pretty, b.tabs), unicode.IsSpace)
	} else {
		str += b.right.Dfl(b.quotes, b.pretty, b.tabs)
	}
	if b.pretty {
		str += "\n"
		str += strings.Repeat(indentSpace, b.indent)
	}
	str += ")"
	return str
}

func (b Builder) Sql() string {
	str := ""
	str += "("
	if b.left != nil {
		str += b.left.Sql(b.pretty, b.tabs)
		str += " "
	}
	str += b.operator
	str += " "
	str += b.right.Sql(b.pretty, b.tabs)
	str += ")"
	return str
}
