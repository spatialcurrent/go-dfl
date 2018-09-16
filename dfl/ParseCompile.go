package dfl

// ParseCompile parses the input expression and compiles the DFL node.
func ParseCompile(in string) (Node, error) {
	n, err := Parse(in)
	if err != nil {
		return n, err
	}
	return n.Compile(), nil
}
