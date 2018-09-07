// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"strings"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-reader/reader"
)

// In is a BinaryOperator that evaluates to true if the left value is in the right value.
// The left value is cast as a string using "fmt.Sprint(lv)".
// If the right value is an array/slice, then evaluated to true if the left value is in the array/slice.
// Otherwise, evaluates to true if the right string is contained by the left string.
type In struct {
	*BinaryOperator
}

func (i In) Dfl(quotes []string, pretty bool, tabs int) string {
	return i.BinaryOperator.Dfl("in", quotes, pretty, tabs)
}

func (i In) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "in",
		"left":  i.Left.Map(),
		"right": i.Right.Map(),
	}
}

func (i In) Compile() Node {
	left := i.Left.Compile()
	right := i.Right.Compile()
	return In{&BinaryOperator{Left: left, Right: right}}
}

func (i In) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {
	vars, lv, err := i.Left.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating in with left value for "+i.Dfl(quotes, false, 0))
	}

	vars, rv, err := i.Right.Evaluate(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, false, errors.Wrap(err, "Error evaluating right value for "+i.Dfl(quotes, false, 0))
	}

	switch lv.(type) {
	case net.IP:
		lv_ip := lv.(net.IP)
		switch rv.(type) {
		case net.IPNet:
			rv_net := rv.(net.IPNet)
			return vars, rv_net.Contains(lv_ip), nil
		case *net.IPNet:
			rv_net := rv.(*net.IPNet)
			return vars, rv_net.Contains(lv_ip), nil
		}
	}

	switch lv.(type) {
	case []string:
		lvs := lv.([]string)
		switch rv.(type) {
		case []string:
			rvs := rv.([]string)
			if len(lvs) == len(rvs) && len(lvs) == 0 {
				return vars, true, nil
			}
			for i, _ := range rvs {
				if rvs[i] == lvs[0] && i+len(lvs) < len(rvs) {
					match := true
					for j, _ := range lvs {
						if rvs[i+j] != lvs[j] {
							match = false
							break
						}
					}
					if match {
						return vars, true, nil
					}
				}
			}
			return vars, false, nil
		}
	case []byte:
		lvb := lv.([]byte)
		switch rv.(type) {
		case []byte:
			rvb := rv.([]byte)
			if len(lvb) == len(rvb) && len(lvb) == 0 {
				return vars, true, nil
			}
			for i, _ := range rvb {
				if rvb[i] == lvb[0] && i+len(lvb) < len(rvb) {
					match := true
					for j, _ := range lvb {
						if rvb[i+j] != lvb[j] {
							match = false
							break
						}
					}
					if match {
						return vars, true, nil
					}
				}
			}
			return vars, false, nil
		case *reader.Cache:
			rvr := rv.(*reader.Cache)
			rvb, err := rvr.ReadAll()
			if err != nil {
				return vars, false, errors.Wrap(err, "error reading all byte for right value in expression "+i.Dfl(quotes, false, 0))
			}
			if len(lvb) == len(rvb) && len(lvb) == 0 {
				return vars, true, nil
			}
			for i, _ := range rvb {
				if rvb[i] == lvb[0] && i+len(lvb) < len(rvb) {
					match := true
					for j, _ := range lvb {
						if rvb[i+j] != lvb[j] {
							match = false
							break
						}
					}
					if match {
						return vars, true, nil
					}
				}
			}
			return vars, false, nil

		}
	}

	lvs := fmt.Sprint(lv)

	switch rv.(type) {
	case Null:
		return vars, false, nil
	case StringSet:
		return vars, rv.(StringSet).Contains(lvs), nil
	case string:
		return vars, strings.Contains(rv.(string), lvs), nil
	case int:
		return vars, strings.Contains(fmt.Sprint(rv), lvs), nil
	case float64:
		return vars, strings.Contains(strconv.FormatFloat(rv.(float64), 'f', 6, 64), lvs), nil
	case map[interface{}]struct{}:
		_, ok := rv.(map[interface{}]struct{})[lv]
		return vars, ok, nil
	case map[string]struct{}:
		_, ok := rv.(map[string]struct{})[lvs]
		return vars, ok, nil
	case map[interface{}]interface{}:
		_, ok := rv.(map[interface{}]interface{})[lvs]
		return vars, ok, nil
	case map[string]interface{}:
		_, ok := rv.(map[string]interface{})[lvs]
		return vars, ok, nil
	case map[string]string:
		_, ok := rv.(map[string]string)[lvs]
		return vars, ok, nil
	case []string:
		for _, x := range rv.([]string) {
			if lvs == x {
				return vars, true, nil
			}
		}
		return vars, false, nil
	case []interface{}:
		for _, x := range rv.([]interface{}) {
			if lvs == fmt.Sprint(x) {
				return vars, true, nil
			}
		}
		return vars, false, nil
	}

	return vars, false, errors.Wrap(err, "Error evaluating in with left value ("+fmt.Sprint(reflect.TypeOf(lv))+") and right value ("+fmt.Sprint(reflect.TypeOf(rv))+")")
}
