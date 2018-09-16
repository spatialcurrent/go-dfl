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
func (ne NotEqual) Dfl(quotes []string, pretty bool, tabs int) string {
	return ne.BinaryOperator.Dfl("!=", quotes, pretty, tabs)
}

// Sql returns the SQL representation of this node as a string
func (ne NotEqual) Sql(pretty bool, tabs int) string {
	switch right := ne.Right.(type) {
	case Literal:
		switch right.Value.(type) {
		case Null:
			if pretty {
				return ne.Left.Sql(pretty, tabs) + " IS NOT NULL"
			}
		}
	}
	return ne.BinaryOperator.Sql("!=", pretty, tabs)
}

func (ne NotEqual) Map() map[string]interface{} {
	return ne.BinaryOperator.Map("notequal", ne.Left, ne.Right)
}

func (ne NotEqual) Compile() Node {
	left := ne.Left.Compile()
	right := ne.Right.Compile()
	return NotEqual{&BinaryOperator{Left: left, Right: right}}
}

func (ne NotEqual) Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) {

	vars, lv, rv, err := ne.EvaluateLeftAndRight(vars, ctx, funcs, quotes)
	if err != nil {
		return vars, 0, err
	}

	switch lv.(type) {
	case []interface{}:
		lv = TryConvertArray(lv.([]interface{}))
	}

	switch lv.(type) {
	case Null:
		switch rv.(type) {
		case Null:
			return vars, false, nil
		}
		return vars, true, nil
	case string:
		lvs := lv.(string)
		switch rv.(type) {
		case Null:
			return vars, true, nil
		case string:
			return vars, lvs != rv.(string), nil
		case int:
			return vars, lvs != fmt.Sprint(rv.(int)), nil
		case uint8:
			return vars, lvs != fmt.Sprint(rv.(uint8)), nil
		case int64:
			return vars, lvs != fmt.Sprint(rv.(int64)), nil
		case float64:
			return vars, lvs != fmt.Sprint(rv.(float64)), nil
		}
	case float64:
		switch rv.(type) {
		case Null:
			return vars, true, nil
		}
	case net.IP:
		lvip := lv.(net.IP)
		switch rv.(type) {
		case Null:
			return vars, true, nil
		case net.IP:
			rvip := rv.(net.IP)
			if len(lvip) != len(rvip) {
				return vars, true, nil
			}
			for i, lvb := range lvip {
				if lvb != rvip[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		}
	case []string:
		lva := lv.([]string)
		switch rv.(type) {
		case Null:
			return vars, true, nil
		case []string:
			rva := rv.([]string)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		}
	case []int:
		lva := lv.([]int)
		switch rv.(type) {
		case Null:
			return vars, true, nil
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if lvs != int(rva[i]) {
					return vars, true, nil
				}
			}
			return vars, false, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		}
	case []uint8:
		lva := lv.([]uint8)
		switch rv.(type) {
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if int(lvs) != rva[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		case []uint8:
			rva := rv.([]uint8)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if float64(lvs) != rva[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		}
	case []float64:
		lva := lv.([]float64)
		switch rv.(type) {
		case Null:
			return vars, true, nil
		case []int:
			rva := rv.([]int)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if lvs != float64(rva[i]) {
					return vars, true, nil
				}
			}
			return vars, false, nil
		case []float64:
			rva := rv.([]float64)
			if len(lva) != len(rva) {
				return vars, true, nil
			}
			for i, lvs := range lva {
				if lvs != rva[i] {
					return vars, true, nil
				}
			}
			return vars, false, nil
		}
	}

	v, err := CompareNumbers(lv, rv)
	if err != nil {
		return vars, 0, err
	}

	return vars, v != 0, err
}
