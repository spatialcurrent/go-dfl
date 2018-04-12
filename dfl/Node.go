package dfl

type Node interface {
	Dfl() string
	Map() map[string]interface{}
}
