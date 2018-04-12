package dfl

type LessThan struct {
	*BinaryOperator
}

func (lt LessThan) Dfl() string {
	return "(" + lt.Left.Dfl() + " < " + lt.Right.Dfl() + ")"
}

func (lt LessThan) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "<",
		"left":  lt.Left.Map(),
		"right": lt.Right.Map(),
	}
}
