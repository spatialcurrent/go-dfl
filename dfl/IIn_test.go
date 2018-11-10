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

func TestIIn(t *testing.T) {

	ctx := map[string]interface{}{
		"a": "cafe",
		"b": []string{"Bar", "Cafe"},
		"c": struct{ foo string }{foo: ""},
	}

	testCases := []TestCase{
		NewTestCase("bar iin @b", ctx, true),
		NewTestCase("Bar iin @b", ctx, true),
		NewTestCase("@a iin @b", ctx, true),
		NewTestCase("bar iin [Bar, Cafe]", ctx, true),
		NewTestCase("bar iin {Bar, Cafe}", ctx, true),
		NewTestCase("fast_food iin [Bar, Cafe]", ctx, false),
		NewTestCase("fast_food iin {Bar, Cafe}", ctx, false),
		NewTestCase("fast_food iin @b", ctx, false),
		NewTestCase("foo iin @c", ctx, true),
		NewTestCase("bar iin @c", ctx, false),
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
