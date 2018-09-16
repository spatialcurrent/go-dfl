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

// Assign is a BinaryOperator which sets the value of the right side to the attribute or variable defined by the left side.
type Assign struct {
	*BinaryOperator
}

func (a Assign) Dfl(quotes []string, pretty bool, tabs int) string {
	if pretty {
		switch a.Left.(type) {
		case *Attribute:
			switch a.Right.(type) {
			case *Function, *Pipe:
				return strings.Repeat("  ", tabs) + "(\n" + a.Left.Dfl(quotes, true, tabs+1) + " := " + strings.TrimLeftFunc(a.Right.Dfl(quotes, pretty, tabs+1), unicode.IsSpace) + "\n" + strings.Repeat("  ", tabs) + ")"
			}
		case *Variable:
			switch a.Right.(type) {
			case *Function, *Pipe:
				return strings.Repeat("  ", tabs) + "(\n" + a.Left.Dfl(quotes, true, tabs+1) + " := " + strings.TrimLeftFunc(a.Right.Dfl(quotes, pretty, tabs+1), unicode.IsSpace) + "\n" + strings.Repeat("  ", tabs) + ")"
			}
		}
		return a.BinaryOperator.Dfl(":= ", quotes, pretty, tabs)
	}
	return "(" + a.Left.Dfl(quotes, pretty, tabs) + " := " + a.Right.Dfl(quotes, pretty, tabs) + ")"
}

func (a Assign) Sql(pretty bool, tabs int) string {
	if pretty {
		switch left := a.Left.(type) {
		case *Variable:
			str := strings.Repeat("  ", tabs) + "WHERE " + a.Right.Sql(pretty, tabs) + "\n"
			str += strings.Repeat("  ", tabs) + "INTO TEMP TABLE " + left.Sql(pretty, tabs) + ";"
			return str
		}
		return ""
	}

	switch left := a.Left.(type) {
	case *Variable:
		return "WHERE " + a.Right.Sql(pretty, tabs) + " INTO TEMP TABLE " + left.Sql(pretty, tabs) + ";"
	}

	return ""

}

func (a Assign) Map() map[string]interface{} {
	return a.BinaryOperator.Map("assign", a.Left, a.Right)
}

// Compile returns a compiled version of this node.
func (a Assign) Compile() Node {
	left := a.Left.Compile()
	right := a.Right.Compile()
	return &Assign{&BinaryOperator{Left: left, Right: right}}
}

func (a Assign) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	switch lva := a.Left.(type) {
	case Attribute:
		vars, rv, err := a.Right.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, rv, errors.Wrap(err, "error processing right value of "+a.Dfl(quotes, false, 0))
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
		vars, rv, err := a.Right.Evaluate(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, rv, errors.Wrap(err, "error processing right value of "+a.Dfl(quotes, false, 0))
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

	return vars, ctx, errors.New("error evaluating declare.  left value (" + a.Left.Dfl(quotes, false, 0) + ") must be an attribute node but is of type " + fmt.Sprint(reflect.TypeOf(a.Left)))
}
