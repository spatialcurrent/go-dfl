// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// dfl.so creates a shared library of Go that can be called by C, C++, or Python
//

package main

import (
	"C"
	"fmt"
	"unsafe"
)

import (
	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

func main() {}

//export Format
func Format(expression *C.char, outputString **C.char) *C.char {

	node, _, err := dfl.Parse(C.GoString(expression))
	if err != nil {
		return C.CString(err.Error())
	}

	*outputString = C.CString(node.Dfl(dfl.DefaultQuotes, false, 0))

	return nil
}

func buildContext(argc C.int, argv **C.char) map[string]interface{} {

	length := int(argc)

	if length == 0 {
		return map[string]interface{}{}
	}

	// Instead of using [1 << 30] using [(1 << 29) - 1] instead
	// as a workaround for 32 bit architectures
	// https://github.com/docker/docker-credential-helpers/pull/61
	// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
	tmpslice := (*[(1 << 29) - 1]*C.char)(unsafe.Pointer(argv))[:length:length]
	vars := make([]string, length)
	for i, s := range tmpslice {
		vars[i] = C.GoString(s)
	}

	ctx := map[string]interface{}{}
	for i := 0; i < len(vars)/2; i++ {
		ctx[vars[i*2]] = dfl.TryConvertString(vars[i*2+1])
	}

	return ctx
}

//export EvaluateBool
func EvaluateBool(exp *C.char, argc C.int, argv **C.char, result *C.int) *C.char {

	fmt.Println("argc", argc, "argv", argv)

	node, err := dfl.ParseCompile(C.GoString(exp))
	if err != nil {
		return C.CString(err.Error())
	}

	vars := map[string]interface{}{}

	_, r, err := dfl.EvaluateBool(node, vars, buildContext(argc, argv), dfl.NewFuntionMapWithDefaults(), dfl.DefaultQuotes)
	if err != nil {
		return C.CString(err.Error())
	}

	if r {
		*result = 1
	} else {
		*result = 0
	}

	return nil
}

//export EvaluateString
func EvaluateString(exp *C.char, argc C.int, argv **C.char, result **C.char) *C.char {

	node, err := dfl.ParseCompile(C.GoString(exp))
	if err != nil {
		return C.CString(err.Error())
	}

	vars := map[string]interface{}{}

	_, r, err := node.Evaluate(vars, buildContext(argc, argv), dfl.NewFuntionMapWithDefaults(), dfl.DefaultQuotes)
	if err != nil {
		return C.CString(err.Error())
	}

	*result = C.CString(fmt.Sprint(r))

	return nil
}
