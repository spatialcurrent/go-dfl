package dfl

type GreaterThan struct {
	*BinaryOperator
}

func (gt GreaterThan) Dfl() string {
  return "("+gt.Left.Dfl() + " > "+ gt.Right.Dfl()+")"
}

func (gt GreaterThan) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": ">",
		"left": gt.Left.Map(),
		"right": gt.Right.Map(),
	}
}
