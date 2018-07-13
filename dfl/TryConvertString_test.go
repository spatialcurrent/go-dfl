// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"net"
	"reflect"
	"testing"
	"time"
)

func TestTryConvertString(t *testing.T) {

	testCases := []struct {
		in   string
		want interface{}
	}{
		{in: "null", want: ""},
		{in: "none", want: ""},
		{in: "", want: ""},
		{in: "true", want: true},
		{in: "false", want: false},
		{in: "1.0", want: 1.0},
		{in: "3", want: 3},
	}

	for _, testCase := range testCases {
		if got := TryConvertString(testCase.in); testCase.want != got {
			t.Errorf("TryConvertString(%q) == %v (%q), want %v (%q)", testCase.in, got, reflect.TypeOf(got), testCase.want, reflect.TypeOf(testCase.want))
		}
	}

}

func TestTryConvertStringTimes(t *testing.T) {

	now := time.Now()

	testCases := []struct {
		in   string
		want interface{}
	}{
		{in: now.Format(time.RFC3339Nano), want: now},
		{in: "2018-05-02T03:28:56Z", want: time.Date(2018, time.May, 2, 3, 28, 56, 0, time.UTC)},
		{in: "2018-01-01", want: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)},
	}

	for _, testCase := range testCases {
		got := TryConvertString(testCase.in)
		want := testCase.want.(time.Time)
		switch got.(type) {
		case time.Time:
			if !got.(time.Time).Equal(want) {
				t.Errorf("TryConvertString(%q) == %v (%q), want %v (%q)", testCase.in, got, reflect.TypeOf(got), testCase.want, reflect.TypeOf(testCase.want))
			}
		default:
			t.Errorf("TryConvertString(%q) == %v (%q), want %v (%q)", testCase.in, got, reflect.TypeOf(got), testCase.want, reflect.TypeOf(testCase.want))
		}
	}

}


func TestTryConvertStringIPv4(t *testing.T) {

	testCases := []struct {
		in   string
		want interface{}
	}{
		{in: "192.168.2.1", want: net.ParseIP("192.168.2.1")},
		{in: "10.10.1.1", want: net.ParseIP("10.10.1.1")},
	}

	for _, testCase := range testCases {
		got := TryConvertString(testCase.in)
		want := testCase.want.(net.IP)
		switch got.(type) {
		case net.IP:
			if !got.(net.IP).Equal(want) {
				t.Errorf("TryConvertString(%q) == %v (%q), want %v (%q)", testCase.in, got, reflect.TypeOf(got), testCase.want, reflect.TypeOf(testCase.want))
			}
		default:
			t.Errorf("TryConvertString(%q) == %v (%q), want %v (%q)", testCase.in, got, reflect.TypeOf(got), testCase.want, reflect.TypeOf(testCase.want))
		}
	}

}
