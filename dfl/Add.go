package dfl

type Add struct {
	*NumericBinaryOperator
}

func (a Add) Dfl() string {
	return "(" + a.Left.Dfl() + " + " + a.Right.Dfl() + ")"
}

func (a Add) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "+",
		"left":  a.Left.Map(),
		"right": a.Right.Map(),
	}
}

func (a Add) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {

	lv, rv, err := a.EvaluateLeftAndRight(ctx, funcs)
	if err != nil {
		return 0, err
	}

	v, err := AddNumbers(lv, rv)
	if err != nil {
		return 0, err
	}

	return v, err
}
