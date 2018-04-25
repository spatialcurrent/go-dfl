package dfl

type LessThan struct {
	*NumericBinaryOperator
}

func (lt LessThan) Dfl() string {
	return "(" + lt.Left.Dfl() + " < " + lt.Right.Dfl() + ")"
}

func (lt LessThan) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "<",
		"left":  lt.Left.Map(),
		"right": lt.Right.Map(),
	}
}

func (lt LessThan) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {

	v, err := lt.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v < 0, nil
}
