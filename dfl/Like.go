package dfl

type Like struct {
	*BinaryOperator
}

func (l Like) Dfl() string {
	return "(" + l.Left.Dfl() + " like " + l.Right.Dfl() + ")"
}

func (l Like) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "like",
		"left":  l.Left.Map(),
		"right": l.Right.Map(),
	}
}
