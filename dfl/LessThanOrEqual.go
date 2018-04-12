package dfl

type LessThanOrEqual struct {
	*BinaryOperator
}

func (lte LessThanOrEqual) Dfl() string {
  return "("+lte.Left.Dfl() + " <= "+ lte.Right.Dfl()+")"
}

func (lte LessThanOrEqual) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "<=",
		"left": lte.Left.Map(),
		"right": lte.Right.Map(),
	}
}
