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

	"github.com/pkg/errors"
)

func TestIn(t *testing.T) {

	ctx := map[string]interface{}{
		"a": "cafe",
		"b": []string{"bar", "cafe"},
		"c": struct{ foo string }{foo: ""},
	}

	testCases := []TestCase{
		NewTestCase("bar in @b", ctx, true),
		NewTestCase("@a in @b", ctx, true),
		NewTestCase("bar in [bar, cafe]", ctx, true),
		NewTestCase("bar in {bar, cafe}", ctx, true),
		NewTestCase("fast_food in [bar, cafe]", ctx, false),
		NewTestCase("fast_food in {bar, cafe}", ctx, false),
		NewTestCase("fast_food in @b", ctx, false),
		NewTestCase("foo in @c", ctx, true),
		NewTestCase("bar in @c", ctx, false),
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
			t.Errorf("TestIn(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
