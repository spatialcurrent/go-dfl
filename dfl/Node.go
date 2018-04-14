package dfl

type Node interface {
	Dfl() string
	Map() map[string]interface{}
	Evaluate(ctx map[string]interface{}, funcs map[string]func(map[string]interface{}, []string) (interface{}, error)) (interface{}, error)
}
