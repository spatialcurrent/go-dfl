package dfl

type GreaterThanOrEqual struct {
	*NumericBinaryOperator
}

func (gte GreaterThanOrEqual) Dfl() string {
	return "(" + gte.Left.Dfl() + " > " + gte.Right.Dfl() + ")"
}

func (gte GreaterThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    ">=",
		"left":  gte.Left.Map(),
		"right": gte.Right.Map(),
	}
}

func (gte GreaterThanOrEqual) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {

	v, err := gte.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v >= 0, nil
}
