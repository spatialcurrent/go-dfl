package dfl

type Cache struct {
	Results map[string]bool
}

func (c *Cache) Has(key string) bool {
	_, ok := c.Results[key]
	return ok
}

func (c *Cache) Get(key string) bool {
	return c.Results[key]
}

func (c *Cache) Set(key string, result bool) {
	c.Results[key] = result
}

func NewCache() *Cache {
	return &Cache{
		Results: map[string]bool{},
	}
}
