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

func TestILike(t *testing.T) {

	ctx := map[string]interface{}{"a": "Cafe and Bar", "b": "bar"}

	testCases := []TestCase{
		NewTestCase("bar ilike cafe", ctx, false),
		NewTestCase("@a ilike cafe", ctx, false),
		NewTestCase("bar ilike @b", ctx, true),
		NewTestCase("@a ilike @b", ctx, false),
		NewTestCase("@a ilike cafe%", ctx, true),
		NewTestCase("@a ilike %bar", ctx, true),
		NewTestCase("@a ilike %and%", ctx, true),
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
			t.Errorf("TestILike(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
