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

func TestCoalesce(t *testing.T) {

	ctx := map[string]interface{}{"a": nil, "b": 10, "c": 10}

	testCases := []TestCase{
		NewTestCase("(@a ?: 10) == 10", ctx, true),
		NewTestCase("(@a ?: @b) == 10", ctx, true),
		NewTestCase("(@a ?: @Z ?: @c) == 10", ctx, true),
		NewTestCase("(@a ?: @Z ?: @c) == 15", ctx, false),
	}

	for _, testCase := range testCases {
		node, err := Parse(testCase.Expression)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error parsing expression \""+testCase.Expression+"\"").Error())
			continue
		}
		node = node.Compile()
		_, got, err := node.Evaluate(map[string]interface{}{}, testCase.Context, NewFuntionMapWithDefaults(), DefaultQuotes)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error evaluating expression \""+testCase.Expression+"\"").Error())
		} else if got != testCase.Result {
			t.Errorf("TestCoalesce(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
