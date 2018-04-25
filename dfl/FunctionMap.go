package dfl

type FunctionMap map[string]func(map[string]interface{}, []string) (interface{}, error)
