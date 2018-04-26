package dfl

type GreaterThan struct {
	*NumericBinaryOperator
}

func (gt GreaterThan) Dfl() string {
	return "(" + gt.Left.Dfl() + " > " + gt.Right.Dfl() + ")"
}

func (gt GreaterThan) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    ">",
		"left":  gt.Left.Map(),
		"right": gt.Right.Map(),
	}
}

func (gt GreaterThan) Compile() Node {
	left := gt.Left.Compile()
	right := gt.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := CompareNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: (v > 0)}
		}
	}
	return GreaterThan{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (gt GreaterThan) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {

	v, err := gt.EvaluateAndCompare(ctx, funcs)
	if err != nil {
		return false, err
	}

	return v > 0, nil
}
