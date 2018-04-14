package dfl

type Attribute struct {
	Name string
}

func (a Attribute) Dfl() string {
	return "@" + a.Name
}

func (a Attribute) Map() map[string]interface{} {
	return map[string]interface{}{
		"attribute": a.Name,
	}
}

func (a Attribute) Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error) {
	if v, ok := ctx[a.Name]; ok {
		return v, nil
	} else {
		return "", nil
	}
}
