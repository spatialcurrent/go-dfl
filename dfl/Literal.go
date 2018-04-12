package dfl

type Literal struct {
	Value string
}

func (l Literal) Dfl() string {
	return "\""+l.Value+"\""
}

func (l Literal) Map() map[string]interface{} {
	return map[string]interface{}{
	  "value": l.Value,
	}
}
