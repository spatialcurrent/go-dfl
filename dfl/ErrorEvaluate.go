package dfl

// ErrorEvaluate is an error returned when an error occurs during evaluation of a Node.
type ErrorEvaluate struct {
	Node   Node     // the name of the Function
	Quotes []string // the quotes to use
}

// Error returns the error as a string.
func (e ErrorEvaluate) Error() string {
	return "error evaluating expression " + e.Node.Dfl(e.Quotes, false, 0)
}
