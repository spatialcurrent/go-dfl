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

func TestAttribute(t *testing.T) {

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
		NewTestCase("(@a?.d ?: 10) == 10", ctx, true),
		NewTestCase("@c?.d == 10", ctx, true),
		NewTestCase("@c.d == 10", ctx, true),
		NewTestCase("@c.e.f == 100", ctx, true),
		NewTestCase("@c.e?.f == 100", ctx, true),
		NewTestCase("@g[0] == 10", ctx, true),
		NewTestCase("@g[0:2] == [10,20]", ctx, true),
		NewTestCase("@h?.i[1:3] == [20,30]", ctx, true),
		NewTestCase("@h.i[1:3] == [20,30]", ctx, true),
		NewTestCase("@j[0:3] == bar", ctx, true),
		NewTestCase("@j[:3] == bar", ctx, true),
		NewTestCase("@k[0]l == m", ctx, true),
		NewTestCase("@k[0]n == o", ctx, false),
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
			t.Errorf("TestAttribute(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
