// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"github.com/pkg/errors"

	"github.com/spatialcurrent/go-dfl/pkg/dfl/syntax"
)

// Item is a Node representing a key-value pair.
type Item struct {
	Key   Node
	Value Node
}

func (i Item) Dfl(quotes []string, pretty bool, tabs int) string {
	return fmt.Sprintf(
		"%v %v %v",
		i.Key.Dfl(quotes, pretty, tabs),
		syntax.DictionarySeparator,
		i.Value.Dfl(quotes, pretty, tabs),
	)
}

func (i Item) Sql(pretty bool, tabs int) string {
	return i.Key.Sql(pretty, tabs) + ": " + i.Value.Sql(pretty, tabs)
}

func (i Item) Map() map[string]interface{} {
	return map[string]interface{}{
		"@type": "item",
		"@value": map[string]interface{}{
			"key":   i.Key.Map(),
			"value": i.Value.Map(),
		},
	}
}

func (i Item) MarshalMap() (interface{}, error) {
	return i.Map(), nil
}

func (i Item) Compile() Node {
	return Item{
		Key:   i.Key.Compile(),
		Value: i.Value.Compile(),
	}
}

func (i Item) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, k, err := i.Key.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return map[string]interface{}{}, nil, errors.Wrap(err, "error evaluating item key")
	}
	vars, v, err := i.Value.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return map[string]interface{}{}, nil, errors.Wrap(err, "error evaluating item value")
	}
	return vars, []interface{}{k, v}, nil
}

func (i Item) Attributes() []string {
	set := make(map[string]struct{})
	for _, x := range i.Key.Attributes() {
		set[x] = struct{}{}
	}
	for _, x := range i.Value.Attributes() {
		set[x] = struct{}{}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}

func (i Item) Variables() []string {
	set := make(map[string]struct{})
	for _, x := range i.Key.Variables() {
		set[x] = struct{}{}
	}
	for _, x := range i.Value.Variables() {
		set[x] = struct{}{}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
