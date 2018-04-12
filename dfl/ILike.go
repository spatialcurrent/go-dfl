package dfl

type ILike struct {
	*BinaryOperator
}

func (i ILike) Dfl() string {
	return "(" + i.Left.Dfl() + " ilike " + i.Right.Dfl() + ")"
}

func (i ILike) Map() map[string]interface{} {
	return map[string]interface{}{
		"op":    "ilike",
		"left":  i.Left.Map(),
		"right": i.Right.Map(),
	}
}
