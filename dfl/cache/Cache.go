// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package cache

import (
	"sync"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-dfl/dfl"
)

// Cache is a struct that stores the compiled nodes to prevent expensive parsing and compiling of DFL expressions.
type Cache struct {
	nodes *sync.Map
}

// Has returns true if the cache has the Node for the given expression in the cache.
func (c *Cache) Has(exp string) bool {
	v, ok := c.nodes.Load(exp)
	if ! ok {
		return false
	}
	n, ok := v.(dfl.Node)
	if ! ok {
		return false
	}
	if n == nil {
		return false
	}
	return true
}

// Get returns the node for the given expression and true, if a valid node is found.
// Otherwise, Get returns (nil, false).
func (c *Cache) Get(exp string) (dfl.Node, bool) {
	v, ok := c.nodes.Load(exp)
	if ! ok {
		return nil, false
	}
	n, ok := v.(dfl.Node)
	if ! ok {
		return nil, false
	}
	if n == nil {
		return nil, false
	}
	return n, true
}

// Set adds the given node to the cache with the expression as key.
func (c *Cache) Set(exp string, node dfl.Node) {
	c.nodes.Store(exp, node)
}

// Parse and compile the DFL expression.
// If the expression was already compiled, use the version found in the cache.
// Otherwise,
func (c *Cache) ParseCompile(exp string) (dfl.Node, error) {
	node, ok := c.Get(exp)
	if ok {
		return node ,nil
	}

	node, err := dfl.ParseCompile(exp)
	if err != nil {
		return node, errors.Wrap(err, "error with parsing expression "+exp)
	}

	c.Set(exp, node)

	return node, nil
}

// MustParseCompile parses and compiles the expression and then sets the resulting node in the cache.
func (c *Cache) MustParseCompile(exp string) dfl.Node {
	node, err := c.ParseCompile(exp)
	if err != nil {
		panic(err)
	}
	return node
}


// NewCache returns a new node cache
func New() *Cache {
	return &Cache{
		nodes: &sync.Map{},
	}
}
