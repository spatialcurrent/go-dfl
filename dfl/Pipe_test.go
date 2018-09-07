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

func TestPipe(t *testing.T) {

	ctx := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"c": 2,
		},
		"d": []string{"bar", "restaurant", "shop"},
	}

	testCases := []TestCase{
		NewTestCase("true | false", ctx, false),
		NewTestCase("@b | @c == 2", ctx, true),
		NewTestCase("@d | bar in @", ctx, true),
		NewTestCase("@d | sort(@,'',true) | limit(@, 1) | @[0] == shop", ctx, true),
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
			t.Errorf("TestPipe(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
