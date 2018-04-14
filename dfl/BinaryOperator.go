package dfl

type BinaryOperator struct {
	Left  Node
	Right Node
}

func (bo BinaryOperator) EvaluateLeftAndRight(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, interface{}, error) {
	lv, err := bo.Left.Evaluate(ctx, funcs)
	if err != nil {
		return false, false, err
	}
	rv, err := bo.Right.Evaluate(ctx, funcs)
	if err != nil {
		return false, false, err
	}
	return lv, rv, nil
}
