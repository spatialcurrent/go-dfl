// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"testing"

	"github.com/pkg/errors"
)

func TestFormatNodes(t *testing.T) {

	ctx := map[string]interface{}{
		"years": 5,
	}

	testCases := []struct {
		Dfl    string
		Ctx    interface{}
		Result string
	}{
		struct {
			Dfl    string
			Ctx    interface{}
			Result string
		}{
			Dfl:    "Hello + ' world, ' +  'I arrived to this planet' + @years +' years ago.'",
			Ctx:    ctx,
			Result: "concat('Hello world, I arrived to this planet', years, ' years ago.')",
		},
	}

	for _, testCase := range testCases {
		node, err := ParseCompile(testCase.Dfl)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error parsing expression \""+testCase.Dfl+"\"").Error())
			continue
		}
		got := node.Sql(false, 0)
		if err != nil {
			t.Errorf(errors.Wrap(err, "Error evaluating expression \""+testCase.Dfl+"\"").Error())
		} else if got != testCase.Result {
			t.Errorf("TestFormatNodes(%q) == %v, want %v", testCase.Dfl, got, testCase.Result)
		}
	}

}
