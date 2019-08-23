// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Package cli is for
//
package cli

import (
	"github.com/spf13/pflag"
)

const (
	FlagUri     = "uri"
	FlagCompile = "compile"
	FlagPretty  = "pretty"
	FlagTabs    = "tabs"

	ShorthandUri     = "u"
	ShorthandCompile = "c"
	ShorthandPretty  = "p"
	ShorthandTabs    = "t"
)

func InitFmtFlags(flag *pflag.FlagSet) {
	flag.StringP(FlagUri, ShorthandUri, "stdin", "uri to DFL file")
	flag.BoolP(FlagCompile, ShorthandCompile, false, "compile expression")
	flag.BoolP(FlagPretty, ShorthandPretty, false, "pretty output")
	flag.IntP(FlagTabs, ShorthandTabs, 0, "tabs")
}
