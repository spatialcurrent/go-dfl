// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-adaptive-functions/af"
	"reflect"
	"strings"
)

// AssignMultiply is a BinaryOperator which sets the multiplied value of the left side and right side to the attribute or variable defined by the left side.
type AssignMultiply struct {
	*BinaryOperator
}

func (a AssignMultiply) Dfl(quotes []string, pretty bool, tabs int) string {
	b := a.Builder("*=", quotes, tabs)
	if pretty {
		b = b.Indent(tabs).Pretty(pretty).Tabs(tabs + 1).TrimRight(pretty)
		switch a.Left.(type) {
		case *Attribute:
			switch a.Right.(type) {
			case *Function, *Pipe:
				return b.Dfl()
			}
		case *Variable:
			switch a.Right.(type) {
			case *Function, *Pipe:
				return b.Dfl()
			}
		}
		return a.BinaryOperator.Dfl("*= ", quotes, pretty, tabs)
	}
	return b.Dfl()
}

func (a AssignMultiply) Sql(pretty bool, tabs int) string {
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
		return "WHERE " + a.Right.Sql(pretty, tabs) + " INTO TEMP TABLE " + left.Sql(pretty, tabs) + ";" // #nosec
	}

	return ""

}

func (a AssignMultiply) Map() map[string]interface{} {
	return a.BinaryOperator.Map("assignmultiply", a.Left, a.Right)
}

// Compile returns a compiled version of this node.
func (a AssignMultiply) Compile() Node {
	left := a.Left.Compile()
	right := a.Right.Compile()
	return &AssignMultiply{&BinaryOperator{Left: left, Right: right}}
}

func (a AssignMultiply) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	switch left := a.Left.(type) {
	case Attribute:
		vars, lv, rv, err := a.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, 0, err
		}
		value, err := af.Multiply.ValidateRun([]interface{}{lv, rv})
		if err != nil {
			return vars, 0, errors.Wrap(err, ErrorEvaluate{Node: a, Quotes: quotes}.Error())
		}
		if len(left.Name) == 0 {
			return vars, value, nil
		}
		if t := reflect.TypeOf(ctx); t.Kind() != reflect.Map {
			ctx = map[string]interface{}{}
		}
		path := left.Name
		obj := ctx
		for len(path) > 0 {
			if !strings.Contains(path, ".") {
				reflect.ValueOf(obj).SetMapIndex(reflect.ValueOf(path), reflect.ValueOf(value))
				break
			}
			pair := strings.SplitN(path, ".", 2)
			objectValue := reflect.ValueOf(obj)
			next := objectValue.MapIndex(reflect.ValueOf(pair[0]))
			if (reflect.TypeOf(next.Interface()).Kind() != reflect.Map) || (!objectValue.IsValid()) || objectValue.IsNil() {
				m := map[string]interface{}{}
				objectValue.SetMapIndex(reflect.ValueOf(pair[0]), reflect.ValueOf(m))
				obj = m
			} else {
				obj = next.Interface()
			}
			path = pair[1]
		}
		return vars, ctx, nil
	case Variable:
		vars, lv, rv, err := a.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
		if err != nil {
			return vars, 0, err
		}
		value, err := af.Multiply.ValidateRun([]interface{}{lv, rv})
		if err != nil {
			return vars, 0, errors.Wrap(err, ErrorEvaluate{Node: a, Quotes: quotes}.Error())
		}
		path := left.Name
		var obj interface{}
		obj = vars
		for len(path) > 0 {
			if !strings.Contains(path, ".") {
				reflect.ValueOf(obj).SetMapIndex(reflect.ValueOf(path), reflect.ValueOf(value))
				break
			}
			pair := strings.SplitN(path, ".", 2)
			objectValue := reflect.ValueOf(obj)
			next := objectValue.MapIndex(reflect.ValueOf(pair[0]))
			if (reflect.TypeOf(next.Interface()).Kind() != reflect.Map) || (!objectValue.IsValid()) || objectValue.IsNil() {
				m := map[string]interface{}{}
				objectValue.SetMapIndex(reflect.ValueOf(pair[0]), reflect.ValueOf(m))
				obj = m
			} else {
				obj = next.Interface()
			}
			path = pair[1]
		}
		return vars, ctx, nil
	}

	return vars, ctx, &ErrorEvaluate{Node: a, Quotes: quotes}
}
