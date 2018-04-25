package dfl

type BinaryOperator struct {
	Left  Node
	Right Node
}

func (bo BinaryOperator) EvaluateLeftAndRight(ctx map[string]interface{}, funcs FunctionMap) (interface{}, interface{}, error) {
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

func (bo BinaryOperator) Attributes() []string {
	set := make(map[string]struct{})
	for _, x := range bo.Left.Attributes() {
		set[x] = struct{}{}
	}
	for _, x := range bo.Right.Attributes() {
		set[x] = struct{}{}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
