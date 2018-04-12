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
