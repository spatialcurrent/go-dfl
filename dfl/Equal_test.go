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

func TestEqual(t *testing.T) {

	ctx := map[string]interface{}{
		"a": 2,
		"b": 3.0,
		"c": "v",
		"d": []int{1, 2, 3, 4},
		"e": []string{"a", "b", "c", "d"},
		"f": []int{137, 80, 78, 71},
	}

	testCases := []TestCase{
		NewTestCase("2 == 7", ctx, false),
		NewTestCase("@a == 1", ctx, false),
		NewTestCase("3.0 == @b", ctx, true),
		NewTestCase("@a == @b", ctx, false),
		NewTestCase("a == b", ctx, false),
		NewTestCase("'a' == 'b'", ctx, false),
		NewTestCase("192.168.2.1 == 192.168.2.1", ctx, true),
		NewTestCase("192.168.2.1 == 192.168.1.1", ctx, false),
		NewTestCase("[1, 2] == [1, 2]", ctx, true),
		NewTestCase("[1, 2] == [1, 3]", ctx, false),
		NewTestCase("[@c, b] == [v, b]", ctx, true),
		NewTestCase("@d == [1, 2, 3, 4]", ctx, true),
		NewTestCase("@d == [1, 2, 3, 4]", ctx, true),
		NewTestCase("@e == [a, b, c, d]", ctx, true),
		NewTestCase("@f == [0x89, 0x50, 0x4E, 0X47]", ctx, true),
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
			t.Errorf("TestEqual(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
