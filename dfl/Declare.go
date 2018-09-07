// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strings"
	"unicode"
)

// Declare is a BinaryOperator which sets the value of the right side to the attribute or variable defined by the left side.
type Declare struct {
	*BinaryOperator
}

func (d Declare) Dfl(quotes []string, pretty bool, tabs int) string {
	if pretty {
		switch d.Left.(type) {
		case *Attribute:
			switch d.Right.(type) {
			case *Function, *Pipe:
				return strings.Repeat("  ", tabs) + "(\n" + d.Left.Dfl(quotes, true, tabs+1) + " := " + strings.TrimLeftFunc(d.Right.Dfl(quotes, pretty, tabs+1), unicode.IsSpace) + "\n" + strings.Repeat("  ", tabs) + ")"
			}
		case *Variable:
			switch d.Right.(type) {
			case *Function, *Pipe:
				return strings.Repeat("  ", tabs) + "(\n" + d.Left.Dfl(quotes, true, tabs+1) + " := " + strings.TrimLeftFunc(d.Right.Dfl(quotes, pretty, tabs+1), unicode.IsSpace) + "\n" + strings.Repeat("  ", tabs) + ")"
			}
		}
		return d.BinaryOperator.Dfl(":= ", quotes, pretty, tabs)
	}
	return "(" + d.Left.Dfl(quotes, pretty, tabs) + " := " + d.Right.Dfl(quotes, pretty, tabs) + ")"
}

func (d Declare) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "declare",
		"left":  d.Left.Map(),
		"right": d.Right.Map(),
	}
}

// Compile returns a compiled version of this node.
func (d Declare) Compile() Node {
	left := d.Left.Compile()
	right := d.Right.Compile()
	return &Declare{&BinaryOperator{Left: left, Right: right}}
}

func (d Declare) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	switch lva := d.Left.(type) {
	case Attribute:
		vars, rv, err := d.Right.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, rv, errors.Wrap(err, "error processing right value of "+d.Dfl(quotes, false, 0))
		}
		if t := reflect.TypeOf(ctx); t.Kind() != reflect.Map {
			ctx = map[string]interface{}{}
		}
		path := lva.Name
		obj := ctx
		for len(path) > 0 {
			if !strings.Contains(path, ".") {
				reflect.ValueOf(obj).SetMapIndex(reflect.ValueOf(path), reflect.ValueOf(rv))
				break
			}
			pair := strings.SplitN(path, ".", 2)
			v := reflect.ValueOf(obj)
			next := v.MapIndex(reflect.ValueOf(pair[0]))
			if (reflect.TypeOf(next.Interface()).Kind() != reflect.Map) || (!v.IsValid()) || v.IsNil() {
				m := map[string]interface{}{}
				v.SetMapIndex(reflect.ValueOf(pair[0]), reflect.ValueOf(m))
				obj = m
			} else {
				obj = next.Interface()
			}
			path = pair[1]
		}
		return vars, ctx, nil
	case Variable:
		vars, rv, err := d.Right.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, rv, errors.Wrap(err, "error processing right value of "+d.Dfl(quotes, false, 0))
		}
		path := lva.Name
		var obj interface{}
		obj = vars
		for len(path) > 0 {
			if !strings.Contains(path, ".") {
				reflect.ValueOf(obj).SetMapIndex(reflect.ValueOf(path), reflect.ValueOf(rv))
				break
			}
			pair := strings.SplitN(path, ".", 2)
			v := reflect.ValueOf(obj)
			next := v.MapIndex(reflect.ValueOf(pair[0]))
			if (reflect.TypeOf(next.Interface()).Kind() != reflect.Map) || (!v.IsValid()) || v.IsNil() {
				m := map[string]interface{}{}
				v.SetMapIndex(reflect.ValueOf(pair[0]), reflect.ValueOf(m))
				obj = m
			} else {
				obj = next.Interface()
			}
			path = pair[1]
		}
		return vars, ctx, nil
	}

	return vars, ctx, errors.New("error evaluating declare.  left value (" + d.Left.Dfl(quotes, false, 0) + ") must be an attribute node but is of type " + fmt.Sprint(reflect.TypeOf(d.Left)))
}
