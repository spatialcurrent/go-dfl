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

func (s Subtract) Compile() Node {
	left := s.Left.Compile()
	right := s.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := SubtractNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: v}
		}
	}
	return Subtract{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
}

func (s Subtract) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {

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