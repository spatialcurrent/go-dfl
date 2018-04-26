package dfl

type Equal struct {
	*NumericBinaryOperator
}

func (e Equal) Dfl() string {
	return "(" + e.Left.Dfl() + " == " + e.Right.Dfl() + ")"
}

func (e Equal) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "equal",
		"left":  e.Left.Map(),
		"right": e.Right.Map(),
	}
}

func (e Equal) Compile() Node {
	return e
}

func (e Equal) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {

	v, err := e.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v == 0, nil
}
