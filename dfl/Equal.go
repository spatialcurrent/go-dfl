package dfl

type Equal struct {
	*BinaryOperator
}

func (l Like) Dfl() string {
	return "(" + l.Left.Dfl() + " == " + l.Right.Dfl() + ")"
}

func (l Like) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "equal",
		"left":  l.Left.Map(),
		"right": l.Right.Map(),
	}
}
