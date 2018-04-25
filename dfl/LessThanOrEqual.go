package dfl

type LessThanOrEqual struct {
	*NumericBinaryOperator
}

func (lte LessThanOrEqual) Dfl() string {
	return "(" + lte.Left.Dfl() + " <= " + lte.Right.Dfl() + ")"
}

func (lte LessThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "<=",
		"left":  lte.Left.Map(),
		"right": lte.Right.Map(),
	}
}

func (lte LessThanOrEqual) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {

	v, err := lte.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v <= 0, nil
}
