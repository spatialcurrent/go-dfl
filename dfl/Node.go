// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfl

import (
	"reflect"
)

// Node is the interface for representing the constructs
// of the Dyanmic Filter Language in an Abstract Syntax Tree.
// This interface is inherited by most structs in the dfl package.
type Node interface {
	Dfl(quotes []string, pretty bool, tabs int) string                                                                                      // returns the DFL expression representation of this node
	Sql(pretty bool, tabs int) string                                                                                                       // returns the SQL representation of this node
	Map() map[string]interface{}                                                                                                            // returns a map representing this node
	Compile() Node                                                                                                                          // compiles this node (and all children).
	Evaluate(vars map[string]interface{}, ctx interface{}, funcs FunctionMap, quotes []string) (map[string]interface{}, interface{}, error) // evaluates the value of a node given a context
	Attributes() []string                                                                                                                   // returns a slice of all attributes used by this node (and all children nodes)
}

// Used with reflection
var nodeType = reflect.TypeOf(new(Node)).Elem()
