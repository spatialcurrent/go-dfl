package dfl

type Counter map[string]int

func (c Counter) Len() int {
	return len(c)
}

func (c Counter) Has(key string) bool {
	_, ok := c[key]
	return ok
}

func (c Counter) Increment(key string) {
	if count, ok := c[key]; ok {
		c[key] = count + 1
	} else {
		c[key] = 1
	}
}
