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

func TestFunction(t *testing.T) {

	ctx := map[string]interface{}{
		"a": 2,
		"b": 3.0,
		"c": "PNG",
		"d": []map[string]interface{}{
			map[string]interface{}{"e": "f"},
			map[string]interface{}{"e": "h"},
		},
	}

	testCases := []TestCase{
		NewTestCase("min(1, @a) == 1", ctx, true),
		NewTestCase("min(2, @a) == 2", ctx, true),
		NewTestCase("min(3.0, @b) == 3.0", ctx, true),
		NewTestCase("min(1, max(@a, 2)) == 1", ctx, true),
		NewTestCase("len(bytes(@c)) == 3", ctx, true),
		NewTestCase("f in map(@d, '@e')", ctx, true),
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
			t.Errorf("TestFunction(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
