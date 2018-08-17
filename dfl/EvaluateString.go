package dfl

import (
	"fmt"
	"reflect"
)

import (
	"github.com/pkg/errors"
)

// EvaluateString returns the string value of a node given a context.  If the result is not a string, then returns an error.
func EvaluateString(n Node, ctx interface{}, funcs FunctionMap) (string, error) {
	result, err := n.Evaluate(ctx, funcs)
	if err != nil {
		return "", errors.Wrap(err, "Error evaluating expression")
	}

	switch result.(type) {
	case string:
		return result.(string), nil
	}

	return "", errors.New("Evaluation returned a " + fmt.Sprint(reflect.TypeOf(result)) + " instead of string")
}
