// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package dfl provides the interfaces, embedded structs, and implementing code
// for parsing, compiling, and evaluating Dynamic Filter Language (DFL) expressions.
//
// Usage
//
// You can import dfl as a package into your own Go project or use the command line interface.
// A common architecture is to have a client application generate a DFL expression string and submit to a Go application using a rest interface.
//
//  import (
//    "github.com/spatialcurrent/go-dfl/dfl"
//  )
//  root, err := dfl.Parse("<YOUR EXPRESSION>")
//  if err != nil {
//    panic(err)
//  }
//  root = root.Compile()  // Flattens tree for performance if possible.
//  result := root.Evaluate(Context{<YOUR KEY:VALUE Context>})
//
// DFL
//
// DFL is another query or filter language like SQL, CQL, or ECQL.  DFL aims to be easily understood by humans in a variety of contexts, such as in a url, in an editor, or in a python terminal.
// The principals are as follows:
//	1.  Easy to Read - The "@" in front of every attribute.
//	2.  Clean - Quotes are optional (unless required because of spaces in an element)
//	3.  Strict Execution Path - Use of parentheses is strongly encouraged o maximize performance over large datasets.
//	4.  Dynamically typed - Operators support multiple types and try to cast if possible.  Fails hard if not valid.
//	5.  Embeddable - Easily written in other languages, such as Python, Javascript, or Shell, without endless worry about escaping.
//
// DFL aims to cover a wide variety of filters while keeping the language expressive and easy to read.  DFL currently supports:
//
//  * Boolean: Not, And, Or
//  * Numeric: LessThan, LessThanOrEqual, Equal, NotEqual, GreaterThan, GreaterThanOrEqual, Add, Subtract
//  * String: Like, ILike, In
//  * Time: Before, After
//  * Array/Set: In
//  * Function: Function
//
// Command Line Interface
//
// See the github.com/go-dfl/cmd/dfl package for a command line tool for testing DFL expressions.
//
//  - https://godoc.org/github.com/spatialcurrent/go-dfl/dfl
//
// Projects
//
// go-dfl is used by the railgun and go-osm project.
//  - https://godoc.org/github.com/spatialcurrent/railgun/railgun
//  - https://godoc.org/github.com/spatialcurrent/go-osm/osm
//  - https://godoc.org/github.com/spatialcurrent/go-dfl/dfl
//
// Examples
//
// Below are some simple examples.
//
//  import (
//    "github.com/spatialcurrent/go-dfl/dfl"
//  )
//  root, err := dfl.Parse("(@amenity in [restaurant, bar]) or (@craft in [brewery, distillery])")
//  if err != nil {
//    panic(err)
//  }
//  root = root.Compile()
//  valid := root.Evaluate(Context{"amenity": "bar", "name": "John's Whiskey Bar"})
//
package dfl

var DefaultQuotes = []string{"'", "\"", "`"}
