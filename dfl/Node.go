package dfl

type Node interface {
	Dfl() string
	Map() map[string]interface{}
	Evaluate(ctx map[string]interface{}, funcs FunctionMap) (interface{}, error)
	Attributes() []string
}
