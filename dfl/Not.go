package dfl

type Not struct {
  *UnaryOperator
}

func (n Not) Dfl() string {
  return "not "+n.Node.Dfl()
}

func (n Not) Map() map[string]interface{} {
	return map[string]interface{}{
	  "op": "not",
		"node": n.Node.Map(),
	}
}
