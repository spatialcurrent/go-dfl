// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Cache is a struct that stores the results of evaluations.
// Cache can be used to speed up evaluations of a complex filter against non-random data, e.g., OSM Tags.
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
