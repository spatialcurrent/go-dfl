// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
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

func (e Equal) Dfl(quotes []string, pretty bool, tabs int) string {
	return e.BinaryOperator.Dfl("==", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (e Equal) Sql(pretty bool, tabs int) string {
	switch right := e.Right.(type) {
	case Literal:
		switch right.Value.(type) {
		case Null:
			if pretty {
				return e.Left.Sql(pretty, tabs) + " IS NULL"
			}
		}
	}
	return e.BinaryOperator.Sql("=", pretty, tabs)
}

func (e Equal) Map() map[string]interface{} {
	return e.BinaryOperator.Map("equal", e.Left, e.Right)
}

func (e Equal) Compile() Node {
	left := e.Left.Compile()
	right := e.Right.Compile()
	return Equal{&BinaryOperator{Left: left, Right: right}}
}

func (e Equal) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, rv, err := e.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	switch lv.(type) {
	case Null, nil:
		switch rv.(type) {
		case Null, nil:
			return vars, true, nil
		}
		return vars, false, nil
	default:
		switch rv.(type) {
		case Null, nil:
			return vars, false, nil
		}
	}

	switch lv.(type) {
	case byte:
		lvs := lv.(byte)
		switch rv.(type) {
		case string:
			return vars, string([]byte{lvs}) == rv.(string), nil
		case uint8:
			return vars, lvs == (rv.(byte)), nil
		}
	case string:
		lvs := lv.(string)
		switch rv.(type) {
		case Null:
			return vars, false, nil
		case string:
			return vars, lvs == rv.(string), nil
		case int:
			return vars, lvs == fmt.Sprint(rv.(int)), nil
		case uint8:
			return vars, lvs == fmt.Sprint(rv.(uint8)), nil
		case int64:
			return vars, lvs == fmt.Sprint(rv.(int64)), nil
		case float64:
			return vars, lvs == fmt.Sprint(rv.(float64)), nil
		}
	case float64:
		switch rv.(type) {
		case Null:
			return vars, false, nil
		}
	case net.IP:
		lvip := lv.(net.IP)
		switch rv.(type) {
		case net.IP:
			rvip := rv.(net.IP)
			if len(lvip) != len(rvip) {
				return vars, false, nil
			}
			for i, lvb := range lvip {
				if lvb != rvip[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		}
	case []string:
		lva := lv.([]string)
		switch rv.(type) {
		case []string:
			rva := rv.([]string)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		}
	case []int:
		lva := lv.([]int)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if lvs != int(rva[i]) {
					return vars, false, nil
				}
			}
			return vars, true, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		}
	case []uint8:
		lva := lv.([]uint8)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if int(lvs) != rva[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		}
	case []float64:
		lva := lv.([]float64)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if lvs != float64(rva[i]) {
					return vars, false, nil
				}
			}
			return vars, true, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return vars, false, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, false, nil
				}
			}
			return vars, true, nil
		}
	}

	v, err := CompareNumbers(lv, rv)
	if err != nil {
		return vars, 0, err
	}

	return vars, v == 0, err
}
