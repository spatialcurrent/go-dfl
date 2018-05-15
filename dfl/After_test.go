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
	"time"
)

import (
	"github.com/pkg/errors"
)

func TestAfter(t *testing.T) {

	ctx := Context{
		"a": time.Date(2018, time.April, 2, 3, 28, 56, 0, time.UTC),
		"b": time.Date(2018, time.May, 2, 3, 28, 56, 0, time.UTC),
	}

	testCases := []TestCase{
		NewTestCase("2017-01-01 after 2018-01-01", ctx, false),
		NewTestCase("@a after 2018-01-01", ctx, true),
		NewTestCase("2017-01-01 after @b", ctx, false),
		NewTestCase("@a after @b", ctx, false),
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
			t.Errorf("TestAfter(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
