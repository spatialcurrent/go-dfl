// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package dfljs

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
	"github.com/spatialcurrent/go-dfl/dfl"
	"honnef.co/go/js/console"
)

// Parse parses a DFL expression and returns a dfljs.Node object that can be used by JavaScript.
func Parse(s string) *js.Object {
	root, _, err := dfl.Parse(s)
	if err != nil {
		console.Error(errors.Wrap(err, "error parsing an expression").Error())
		return js.MakeWrapper(Node{Node: nil})
	}
	return js.MakeWrapper(Node{Node: root})
}
