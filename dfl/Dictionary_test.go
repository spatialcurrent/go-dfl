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
)

import (
	"github.com/pkg/errors"
)

func TestDictionary(t *testing.T) {

	ctx := map[string]interface{}{"a": 2, "b": 3.0}

	testCases := []TestCase{
		NewTestCase("{c: @a}", ctx, map[interface{}]interface{}{"c": 2}),
		NewTestCase("{'c': @}", ctx, map[interface{}]interface{}{"c": ctx}),
		NewTestCase("{'c': {d:@a}}", ctx, map[interface{}]interface{}{"c": map[interface{}]interface{}{"d": 2}}),
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
		} else if !reflect.DeepEqual(got, testCase.Result) {
			t.Errorf("TestDictionary(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
