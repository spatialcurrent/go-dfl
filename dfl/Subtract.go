package dfl

type Subtract struct {
	*NumericBinaryOperator
}

func (s Subtract) Dfl() string {
	return "(" + s.Left.Dfl() + " - " + s.Right.Dfl() + ")"
}

func (s Subtract) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "-",
		"left":  s.Left.Map(),
		"right": s.Right.Map(),
	}
}

func (s Subtract) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {

	lv, rv, err := s.EvaluateLeftAndRight(ctx, funcs)
	if err != nil {
		return 0, err
	}

	v, err := SubtractNumbers(lv, rv)
	if err != nil {
		return 0, err
	}

	return v, err
}
