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
	FlagOutputGo   = "go"
	FlagOutputJSON = "json"
	FlagOutputYAML = "yaml"

	ShorthandOutputGo   = "g"
	ShorthandOutputJSON = "j"
	ShorthandOutputYAML = "y"
)

func InitOutputFlags(flag *pflag.FlagSet) {
	flag.BoolP(FlagOutputGo, ShorthandOutputGo, false, "go output")
	flag.BoolP(FlagOutputJSON, ShorthandOutputJSON, false, "json output")
	flag.BoolP(FlagOutputYAML, ShorthandOutputYAML, false, "yaml output")
}
