// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"strings"
)

// Dictionary is a Node representing a dictionary of key value pairs.
type Dictionary struct {
	//Nodes map[Node]Node
	Items []Item
}

func NewDictionary(m map[string]interface{}) *Dictionary {
	items := make([]Item, 0)
	for k, v := range m {
		if d, ok := v.(map[string]interface{}); ok {
			items = append(items, Item{
				Key:   &Literal{Value: k},
				Value: NewDictionary(d),
			})
		} else {
			items = append(items, Item{
				Key:   &Literal{Value: k},
				Value: &Literal{Value: v},
			})
		}
	}
	return &Dictionary{Items: items}
}

// Len returns the length of the underlying array.
func (d Dictionary) Len() int {
	return len(d.Items)
}

func (d Dictionary) Dfl(quotes []string, pretty bool, tabs int) string {
	if len(d.Items) == 0 {
		return "{}"
	}
	values := make([]string, 0)
	for _, i := range d.Items {
		values = append(values, i.Key.Dfl(quotes, pretty, tabs+1)+": "+i.Value.Dfl(quotes, pretty, tabs+1))
	}
	if pretty {
		return "{" + "\n" + FormatList(values, ",", pretty, tabs+1) + "\n" + strings.Repeat(DefaultTab, tabs) + "}"
	}
	return "{" + FormatList(values, ",", pretty, tabs) + "}"

}

// Sql returns the SQL representation of this node as a string
func (d Dictionary) Sql(pretty bool, tabs int) string {
	str := SqlQuote + SqlArrayPrefix
	for i, item := range d.Items {
		if i > 0 {
			str += ", "
		}
		str += item.Key.Sql(pretty, tabs) + ":" + item.Value.Sql(pretty, tabs)
		i += 1
	}
	str = str + SqlArraySuffix + SqlQuote + "::json"
	return str
}

func (d Dictionary) Map() map[string]interface{} {
	items := []map[string]interface{}{}
	for _, item := range d.Items {
		items = append(items, item.Map())
	}
	return map[string]interface{}{
		"@type":  "dictionary",
		"@value": items,
	}
}

func (d Dictionary) MarshalMap() (interface{}, error) {
	return d.Map(), nil
}

// Compile returns a compiled version of this node.
func (d Dictionary) Compile() Node {
	items := make([]Item, 0)
	for _, i := range d.Items {
		items = append(items, Item{
			Key:   i.Key.Compile(),
			Value: i.Value.Compile(),
		})
	}
	return &Dictionary{Items: items}
}

func (d Dictionary) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	values := map[interface{}]interface{}{}
	for _, i := range d.Items {
		_, keyValue, err := i.Key.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, values, err
		}
		_, valueValue, err := i.Value.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, values, err
		}
		values[keyValue] = valueValue
	}
	return vars, values, nil
}

func (d Dictionary) Attributes() []string {
	set := make(map[string]struct{})
	for _, i := range d.Items {
		for _, x := range i.Key.Attributes() {
			set[x] = struct{}{}
		}
		for _, x := range i.Value.Attributes() {
			set[x] = struct{}{}
		}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}

func (d Dictionary) Variables() []string {
	set := make(map[string]struct{})
	for _, i := range d.Items {
		for _, x := range i.Key.Variables() {
			set[x] = struct{}{}
		}
		for _, x := range i.Value.Variables() {
			set[x] = struct{}{}
		}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
