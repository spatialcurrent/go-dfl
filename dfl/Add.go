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

func (a Add) Compile() Node {
	left := a.Left.Compile()
	right := a.Right.Compile()
	switch left.(type) {
	case Literal:
		switch right.(type) {
		case Literal:
			v, err := AddNumbers(left.(Literal).Value, right.(Literal).Value)
			if err != nil {
				panic(err)
			}
			return Literal{Value: v}
		}
	}
	return Add{&NumericBinaryOperator{&BinaryOperator{Left: left, Right: right}}}
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
