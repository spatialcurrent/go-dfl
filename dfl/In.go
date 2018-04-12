package dfl

type In struct {
	*BinaryOperator
}

func (i In) Dfl() string {
	return "(" + i.Left.Dfl() + " in " + i.Right.Dfl() + ")"
}

func (i In) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "in",
		"left":  i.Left.Map(),
		"right": i.Right.Map(),
	}
}
