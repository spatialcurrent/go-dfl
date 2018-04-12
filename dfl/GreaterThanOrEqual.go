package dfl

type GreaterThanOrEqual struct {
	*BinaryOperator
}

func (gte GreaterThanOrEqual) Dfl() string {
  return "("+gte.Left.Dfl() + " > "+ gte.Right.Dfl()+")"
}

func (gte GreaterThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": ">=",
		"left": gte.Left.Map(),
		"right": gte.Right.Map(),
	}
}
