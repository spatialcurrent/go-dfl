package dfl

type Or struct {
	*BinaryOperator
}

func (o Or) Dfl() string {
	return "(" + o.Left.Dfl() + " or " + o.Right.Dfl() + ")"
}

func (o Or) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "or",
		"left":  o.Left.Map(),
		"right": o.Right.Map(),
	}
}
