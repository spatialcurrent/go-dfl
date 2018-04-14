package dfl

import (
	"github.com/pkg/errors"
)

type Not struct {
	*UnaryOperator
}

func (n Not) Dfl() string {
	return "not " + n.Node.Dfl()
}

func (n Not) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":   "not",
		"node": n.Node.Map(),
	}
}

func (n Not) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {
	v, err := n.Node.Evaluate(ctx, funcs)
	if err != nil {
		return false, err
	}
	switch v.(type) {
	case bool:
		return !(v.(bool)), nil
	}
	return false, errors.New("Error evaulating expression " + n.Dfl())
}
