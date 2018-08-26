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
)

// Equal is a BinaryOperator that evaluating to true if parameter a is equal to parameter b.
// The parameters may be of type int, int64, or float64.
type Equal struct {
	*BinaryOperator
}

func (e Equal) Dfl(quotes []string, pretty bool) string {
	return "(" + e.Left.Dfl(quotes, pretty) + " == " + e.Right.Dfl(quotes, pretty) + ")"
}

func (e Equal) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "equal",
		"left":  e.Left.Map(),
		"right": e.Right.Map(),
	}
}

func (e Equal) Compile() Node {
	left := e.Left.Compile()
	right := e.Right.Compile()
	return Equal{&BinaryOperator{Left: left, Right: right}}
}

func (e Equal) Evaluate(ctx interface{}, funcs FunctionMap, quotes []string) (interface{}, error) {

	lv, rv, err := e.EvaluateLeftAndRight(ctx, funcs, quotes)
	if err != nil {
		return 0, err
	}

	switch lv.(type) {
	case Null:
		switch rv.(type) {
		case Null:
			return true, nil
		}
		return false, nil
	case byte:
		lvs := lv.(byte)
		switch rv.(type) {
		case string:
			return string([]byte{lvs}) == rv.(string), nil
		case uint8:
			return lvs == (rv.(byte)), nil
		}
	case string:
		lvs := lv.(string)
		switch rv.(type) {
		case Null:
			return false, nil
		case string:
			return lvs == rv.(string), nil
		case int:
			return lvs == fmt.Sprint(rv.(int)), nil
		case uint8:
			return lvs == fmt.Sprint(rv.(uint8)), nil
		case int64:
			return lvs == fmt.Sprint(rv.(int64)), nil
		case float64:
			return lvs == fmt.Sprint(rv.(float64)), nil
		}
	case float64:
		switch rv.(type) {
		case Null:
			return false, nil
		}
	case net.IP:
		lvip := lv.(net.IP)
		switch rv.(type) {
		case net.IP:
			rvip := rv.(net.IP)
			if len(lvip) != len(rvip) {
				return false, nil
			}
			for i, lvb := range lvip {
				if lvb != rvip[i] {
					return false, nil
				}
			}
			return true, nil
		}
	case []string:
		lva := lv.([]string)
		switch rv.(type) {
		case []string:
			rva := rv.([]string)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return false, nil
				}
			}
			return true, nil
		}
	case []int:
		lva := lv.([]int)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return false, nil
				}
			}
			return true, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if lvs != int(rva[i]) {
					return false, nil
				}
			}
			return true, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return false, nil
				}
			}
			return true, nil
		}
	case []uint8:
		lva := lv.([]uint8)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if int(lvs) != rva[i] {
					return false, nil
				}
			}
			return true, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return false, nil
				}
			}
			return true, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return false, nil
				}
			}
			return true, nil
		}
	case []float64:
		lva := lv.([]float64)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if lvs != float64(rva[i]) {
					return false, nil
				}
			}
			return true, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return false, nil
				}
			}
			return true, nil
		}
	}

	v, err := CompareNumbers(lv, rv)
	if err != nil {
		return 0, err
	}

	return v == 0, err
}
