// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

// Context is a simple alias for a map[string]interface{} that is used for containing the context for evaluating a DFL Node.
// The values in a context are essentially the input parameters for a DFL expression and match up with the Attribute.
// The Context is built from the trailing command line arguments.  For example the arguments from the following command line
//
//  ./dfl -filter "(@amenity like bar) and (open > 0)" amenity=bar popularity=10 open=1
//
// Would be interpreted as the following Context
//
//  ctx := Context{"amenity": "bar", "popularity": 10, "open": 1}
//
type Context struct {
	Data map[string]interface{}
}

func (c *Context) Has(key string) bool {
	_, ok := c.Data[key]
	return ok
}

func (c *Context) Get(key string) interface{} {
	return c.Data[key]
}

func (c *Context) Set(key string, value interface{}) {
	c.Data[key] = value
}
