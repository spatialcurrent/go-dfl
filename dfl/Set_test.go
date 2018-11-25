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

func TestSet(t *testing.T) {

	ctx := map[string]interface{}{}

	testCases := []TestCase{
		NewTestCase("{x}", ctx, map[string]struct{}{"x": struct{}{}}),
		NewTestCase("{x,y}", ctx, map[string]struct{}{"x": struct{}{}, "y": struct{}{}}),
		NewTestCase("{x,y,z}", ctx, map[string]struct{}{"x": struct{}{}, "y": struct{}{}, "z": struct{}{}}),
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
			t.Errorf("TestSet(%q) == %v (%q), want %v (%q)", testCase.Expression, got, reflect.TypeOf(got), testCase.Result, reflect.TypeOf(testCase.Result))
		}
	}

}
