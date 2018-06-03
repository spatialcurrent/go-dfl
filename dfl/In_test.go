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

func TestIn(t *testing.T) {

	ctx := Context{"a": "cafe", "b": []string{"bar", "cafe"}}

	testCases := []TestCase{
		NewTestCase("bar in @b", ctx, true),
		NewTestCase("@a in @b", ctx, true),
		NewTestCase("bar in [bar, cafe]", ctx, true),
		NewTestCase("bar in {bar, cafe}", ctx, true),
		NewTestCase("fast_food in [bar, cafe]", ctx, false),
		NewTestCase("fast_food in {bar, cafe}", ctx, false),
		NewTestCase("fast_food in @b", ctx, false),
	}

	for _, testCase := range testCases {
		node, err := Parse(testCase.Expression)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error parsing expression \""+testCase.Expression+"\"").Error())
			continue
		}
		node = node.Compile()
		got, err := node.Evaluate(testCase.Context, FunctionMap{})
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error evaluating expression \""+testCase.Expression+"\"").Error())
		} else if got != testCase.Result {
			t.Errorf("TestILike(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}