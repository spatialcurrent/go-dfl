// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
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

func TestTernaryOperator(t *testing.T) {

	ctx := map[string]interface{}{
		"a": nil,
		"b": nil,
		"c": map[string]interface{}{"d": 10, "e": map[string]interface{}{"f": 100}},
		"g": []int{10, 20, 30},
		"h": map[string]interface{}{"i": []int{10, 20, 30, 40}},
		"j": "bars",
		"k": []map[string]interface{}{
			map[string]interface{}{"l": "m"},
		},
	}

	testCases := []TestCase{
		NewTestCase("true ? 1 : 2", ctx, 1),
		NewTestCase("false ? 1 : 2", ctx, 2),
		NewTestCase("((@a == null) ? 10 : 20) == 10", ctx, true),
		NewTestCase("((@a == null) ? 10 : 20) | @ += 30", ctx, 40),
		NewTestCase("((@a != null) ? 10 : 50) | @ - 20", ctx, 30),
		NewTestCase(
			"map(@, '(@name == foo) ? (@ + {name: bar}) : @')",
			[]interface{}{map[string]interface{}{"name": "foo"}},
			[]interface{}{map[interface{}]interface{}{"name": "bar"}},
		),
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
		} else if !reflect.DeepEqual(got, testCase.Result) {
			t.Errorf("TestAttribute(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
