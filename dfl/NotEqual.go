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

// NotEqual is a BinaryOperator that evaluates to true if the left value is not equal to the right value.
// The values may be of type int, int64, or float64.
type NotEqual struct {
	*BinaryOperator
}

// Dfl returns the DFL expression representation of the node as a string value.
// For example
//	"( @amenity  !=  shop )"
func (ne NotEqual) Dfl() string {
	return "(" + ne.Left.Dfl() + " != " + ne.Right.Dfl() + ")"
}

func (ne NotEqual) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "equal",
		"left":  ne.Left.Map(),
		"right": ne.Right.Map(),
	}
}

func (ne NotEqual) Compile() Node {
	left := ne.Left.Compile()
	right := ne.Right.Compile()
	return NotEqual{&BinaryOperator{Left: left, Right: right}}
}

func (ne NotEqual) Evaluate(ctx interface{}, funcs FunctionMap) (interface{}, error) {

	lv, rv, err := ne.EvaluateLeftAndRight(ctx, funcs)
	if err != nil {
		return 0, err
	}

	switch lv.(type) {
	case Null:
		switch rv.(type) {
		case Null:
			return false, nil
		}
		return true, nil
	case string:
		lvs := lv.(string)
		switch rv.(type) {
		case Null:
			return true, nil
		case string:
			return lvs != rv.(string), nil
		case int:
			return lvs != fmt.Sprint(rv.(int)), nil
		case uint8:
			return lvs != fmt.Sprint(rv.(uint8)), nil
		case int64:
			return lvs != fmt.Sprint(rv.(int64)), nil
		case float64:
			return lvs != fmt.Sprint(rv.(float64)), nil
		}
	case float64:
		switch rv.(type) {
		case Null:
			return true, nil
		}
	case net.IP:
		lvip := lv.(net.IP)
		switch rv.(type) {
		case Null:
			return true, nil
		case net.IP:
			rvip := rv.(net.IP)
			if len(lvip) != len(rvip) {
				return true, nil
			}
			for i, lvb := range lvip {
				if lvb != rvip[i] {
					return true, nil
				}
			}
			return false, nil
		}
	case []string:
		lva := lv.([]string)
		switch rv.(type) {
		case Null:
			return true, nil
		case []string:
			rva := rv.([]string)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return true, nil
				}
			}
			return false, nil
		}
	case []int:
		lva := lv.([]int)
		switch rv.(type) {
		case Null:
			return true, nil
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return true, nil
				}
			}
			return false, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if lvs != int(rva[i]) {
					return true, nil
				}
			}
			return false, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return true, nil
				}
			}
			return false, nil
		}
	case []uint8:
		lva := lv.([]uint8)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if int(lvs) != rva[i] {
					return true, nil
				}
			}
			return false, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return true, nil
				}
			}
			return false, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return true, nil
				}
			}
			return false, nil
		}
	case []float64:
		lva := lv.([]float64)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if lvs != float64(rva[i]) {
					return true, nil
				}
			}
			return false, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return true, nil
				}
			}
			return false, nil
		}
	}

	v, err := CompareNumbers(lv, rv)
	if err != nil {
		return 0, err
	}

	return v != 0, err
}
