package dfl

import (
	"fmt"
)

import (
	"github.com/pkg/errors"
)

func SubtractNumbers(a interface{}, b interface{}) (interface{}, error) {
	switch a.(type) {
	case int:
		switch b.(type) {
		case int:
			return a.(int) - b.(int), nil
		case int64:
			return int64(a.(int)) - b.(int64), nil
		case float64:
			return float64(a.(int)) - b.(float64), nil
		}
	case int64:
		switch b.(type) {
		case int:
			return a.(int64) - int64(b.(int)), nil
		case int64:
			return a.(int64) - b.(int64), nil
		case float64:
			return float64(a.(int64)) - b.(float64), nil
		}
	case float64:
		switch b.(type) {
		case int:
			return a.(float64) - float64(b.(int)), nil
		case int64:
			return a.(float64) - float64(b.(int64)), nil
		case float64:
			return a.(float64) - b.(float64), nil
		}
	}

	return 0, errors.New("Error subtracting " + fmt.Sprint(a) + " - " + fmt.Sprint(b))
}
