package dfl

import (
	"fmt"
	"reflect"
)

import (
	"github.com/pkg/errors"
)

// EvaluateInt returns the int value of a node given a context.  If the result is not an int, then returns an error.
func EvaluateInt(n Node, ctx Context, funcs FunctionMap) (int, error) {
	result, err := n.Evaluate(ctx, funcs)
	if err != nil {
		return 0, errors.Wrap(err, "Error evaluating expression")
	}

	switch result.(type) {
	case int:
		return result.(int), nil
	case float64:
		return int(result.(float64)), nil
	}

	return 0, errors.New("Evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of int")
}
