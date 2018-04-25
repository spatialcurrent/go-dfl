package dfl

type NotEqual struct {
	*NumericBinaryOperator
}

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

func (ne NotEqual) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {

	v, err := ne.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v != 0, nil
}
