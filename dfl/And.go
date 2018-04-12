package dfl

type And struct {
  *BinaryOperator
}

func (a And) Dfl() string {
  return "("+a.Left.Dfl() + " and " + a.Right.Dfl()+")"
}

func (a And) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "and",
		"left": a.Left.Map(),
		"right": a.Right.Map(),
	}
}
