package dfl

type Array struct {
	Nodes []Node
}

func (a Array) Dfl() string {
	str := "["
	for i, x := range a.Nodes {
		if i > 0 {
			str += ", "
		}
		str += x.Dfl()
	}
	str = str + "]"
	return str
}

func (a Array) Map() map[string]interface{} {
	return map[string]interface{}{
		"nodes": a.Nodes,
	}
}
func (a Array) Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error) {
	values := make([]interface{}, len(a.Nodes))
	for i, n := range a.Nodes {
		v, err := n.Evaluate(ctx, funcs)
		if err != nil {
			return values, err
		}
		values[i] = v
	}
	return values, nil
}

func (a Array) Attributes() []string {
	set := make(map[string]struct{})
	for _, n := range a.Nodes {
		for _, x := range n.Attributes() {
			set[x] = struct{}{}
		}
	}
	attrs := make([]string, 0, len(set))
	for x := range set {
		attrs = append(attrs, x)
	}
	return attrs
}
