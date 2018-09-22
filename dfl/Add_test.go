// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
	"testing"
)

import (
	"github.com/pkg/errors"
)

func TestAdd(t *testing.T) {

	ctx := map[string]interface{}{"a": 2, "b": 3.0}

	testCases := []TestCase{
		NewTestCase("2 + 7", ctx, 9),
		NewTestCase("@a + 1", ctx, 3),
		NewTestCase("3.0 + @b", ctx, 6.0),
		NewTestCase("@a + @b", ctx, 5.0),
	}

	for _, testCase := range testCases {
		node, err := ParseCompile(testCase.Expression)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error parsing expression \""+testCase.Expression+"\"").Error())
			continue
		}
		_, got, err := node.Evaluate(map[string]interface{}{}, testCase.Context, NewFuntionMapWithDefaults(), DefaultQuotes)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error evaluating expression \""+testCase.Expression+"\"").Error())
		} else if got != testCase.Result {
			t.Errorf("TestAdd(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}

func TestAddSql(t *testing.T) {

	testCases := []struct {
		Dfl string
		Sql string
	}{
		struct {
			Dfl string
			Sql string
		}{Dfl: "@population + 2", Sql: "(population + 2)"},
		struct {
			Dfl string
			Sql string
		}{Dfl: "@name + '-' + @type", Sql: "concat(name, '-', type)"},
	}

	for _, testCase := range testCases {
		node, err := ParseCompile(testCase.Dfl)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error parsing expression \""+testCase.Dfl+"\"").Error())
			continue
		}
		got := node.Sql(false, 0)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error evaluating expression \""+testCase.Dfl+"\"").Error())
		} else if got != testCase.Sql {
			t.Errorf("TestAdd(%q) == %v , want %v", testCase.Dfl, got, testCase.Sql)
		}
	}

}
