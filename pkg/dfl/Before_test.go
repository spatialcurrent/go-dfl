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
	"time"

	"github.com/pkg/errors"
)

func TestBefore(t *testing.T) {

	ctx := map[string]interface{}{
		"a": time.Date(2018, time.April, 2, 3, 28, 56, 0, time.UTC),
		"b": time.Date(2018, time.May, 2, 3, 28, 56, 0, time.UTC),
	}

	testCases := []TestCase{
		NewTestCase("2017-01-01 before 2018-01-01", ctx, true),
		NewTestCase("@a before 2018-01-01", ctx, false),
		NewTestCase("2017-01-01 before @b", ctx, true),
		NewTestCase("@a before @b", ctx, true),
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
			t.Errorf("TestBefore(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
