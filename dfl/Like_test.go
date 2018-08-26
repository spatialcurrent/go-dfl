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

func TestLike(t *testing.T) {

	ctx := map[string]interface{}{"a": "cafe and bar", "b": "bar"}

	testCases := []TestCase{
		NewTestCase("bar like cafe", ctx, false),
		NewTestCase("@a like cafe", ctx, false),
		NewTestCase("bar like @b", ctx, true),
		NewTestCase("@a like @b", ctx, false),
		NewTestCase("@a like cafe%", ctx, true),
		NewTestCase("@a like %bar", ctx, true),
		NewTestCase("@a like %and%", ctx, true),
	}

	for _, testCase := range testCases {
		node, err := Parse(testCase.Expression)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error parsing expression \""+testCase.Expression+"\"").Error())
			continue
		}
		node = node.Compile()
		got, err := node.Evaluate(testCase.Context, NewFuntionMapWithDefaults(), DefaultQuotes)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error evaluating expression \""+testCase.Expression+"\"").Error())
		} else if got != testCase.Result {
			t.Errorf("TestLike(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
