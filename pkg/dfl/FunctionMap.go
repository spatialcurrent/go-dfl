// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/russross/blackfriday.v2"

	"github.com/spatialcurrent/go-adaptive-functions/pkg/af"
	"github.com/spatialcurrent/go-reader-writer/pkg/io"
	"github.com/spatialcurrent/go-simple-serializer/pkg/gss"
)

// FunctionMap is a map of functions by string that are reference by name in the Function Node.
type FunctionMap map[string]func(FunctionMap, map[string]interface{}, interface{}, []interface{}, []string) (interface{}, error)

func NewFuntionMapWithDefaults() FunctionMap {
	funcs := FunctionMap{}

	for _, fn := range af.Functions {
		for _, alias := range fn.Aliases {
			funcs[alias] = func(fn af.Function) func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
				return func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
					out, err := fn.ValidateRun(args...)
					if err != nil {
						return Null{}, errors.Wrap(err, "Invalid arguments")
					}
					return out, nil
				}
			}(fn)
		}
	}

	funcs["parse"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) != 2 {
			return 0, errors.New("Invalid number of arguments to parse.")
		}

		str, ok := args[0].(string)
		if !ok {
			return 0, errors.New("Invalid arguments to parse.")
		}

		f, ok := args[1].(string)
		if !ok {
			return 0, errors.New("Invalid arguments to parse.")
		}

		t, err := gss.GetType([]byte(str), f)
		if err != nil {
			return "", errors.Wrap(err, "error creating new object for format "+f)
		}

		return gss.DeserializeBytes(&gss.DeserializeBytesInput{
			Bytes:     []byte(str),
			Type:      t,
			Format:    f,
			Header:    []interface{}{},
			Comment:   "",
			SkipLines: 0,
			Limit:     -1,
		})
	}

	funcs["md"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to md.")
		}

		switch s := args[0].(type) {
		case string:
			unsafe := strings.TrimSpace(string(blackfriday.Run(
				[]byte(s),
				blackfriday.WithExtensions(blackfriday.NoExtensions),
				blackfriday.WithRenderer(blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
					Flags: blackfriday.HTMLFlagsNone,
				})),
			)))
			if strings.HasPrefix(unsafe, "<p>") && strings.HasSuffix(unsafe, "</p>") {
				unsafe = unsafe[3 : len(unsafe)-4]
			}
			return unsafe, nil
		}
		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["filter"] = filterArray
	funcs["group"] = groupArray
	funcs["map"] = mapArray
	//funcs["sort"] = sortArray
	funcs["dict"] = toDict
	funcs["hist"] = histArray
	funcs["prefix"] = prefix
	funcs["suffix"] = suffix
	funcs["trim"] = trimString
	funcs["ltrim"] = trimStringLeft
	funcs["rtrim"] = trimStringRight

	funcs["slugify"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
		if len(args) < 2 {
			return 0, errors.New("Invalid number of arguments to slugify.")
		}

		switch s := args[0].(type) {
		case string:
			switch replacement := args[1].(type) {
			case string:
				reg, err := regexp.Compile("[^a-zA-Z0-9]+")
				if err != nil {
					return Null{}, errors.Wrap(err, "Invalid regular expression ")
				}
				return reg.ReplaceAllString(strings.ToLower(s), replacement), nil
			}
		}

		return Null{}, errors.New("Invalid argument of type " + reflect.TypeOf(args[0]).String())
	}

	funcs["first"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

		if v, ok := args[0].(io.ByteReadCloser); ok {
			return v.ReadFirst()
		}

		return af.First.ValidateRun(args)

	}

	funcs["last"] = func(funcs FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {

		if v, ok := args[0].(io.ByteReadCloser); ok {
			b, err := v.ReadAll()
			if err != nil {
				return byte(0), err
			}
			if len(b) == 0 {
				return byte(0), errors.New("last cannot run on an empty io.ByteReadCloser")
			}
			return b[len(b)-1], nil
		}

		return af.Last.ValidateRun(args)
	}

	return funcs
}

var DefaultFunctionMap = NewFuntionMapWithDefaults()
