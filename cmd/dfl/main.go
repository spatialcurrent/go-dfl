// =================================================================
//
// Copyright (C) 2018 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// Dfl is the command line utility for testing DFL expressions.
//
// Usage
//
// Pass your DFL expression using the -filter keyword argument and your context with trailing K=V arguments.
// go-dfl will attempt to convert string values into their approriate types using TryConvertString.
//
//	Usage: dfl -filter INPUT [-verbose] [-version] [-help] [A=1] [B=2]
//	Options:
//		-filter string
//			The DFL expression to evaulate
//		-help
//			Print help
//		-verbose
//			Provide verbose output
//		-version
//			Prints version to stdout
//
package main

import (
	//"bufio"
	//"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

import (
	"github.com/spatialcurrent/go-dfl/dfl"
)

var GO_DFL_VERSION = "0.0.3"

func main() {

	start := time.Now()

	var filter_text string

	var load_env bool
	var verbose bool
	var version bool
	var help bool

	flag.StringVar(&filter_text, "f", "", "The DFL expression to evaulate")

	flag.BoolVar(&load_env, "env", false, "Load environment variables")
	flag.BoolVar(&verbose, "verbose", false, "Provide verbose output")
	flag.BoolVar(&version, "version", false, "Prints version to stdout")
	flag.BoolVar(&help, "help", false, "Print help")

	flag.Parse()

	if help {
		fmt.Println("Usage: dfl -f INPUT [-verbose] [-version] [-help] [-env] [A=1] [B=2]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: Provided no arguments.")
		fmt.Println("Run \"dfl -help\" for more information.")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Println("Usage: dfl -f INPUT [-verbose] [-version] [-help] [-env] [A=1] [B=2]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println(GO_DFL_VERSION)
		os.Exit(0)
	}

	ctx := map[string]interface{}{}

	if load_env {
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			ctx[strings.TrimSpace(pair[0])] = dfl.TryConvertString(strings.TrimSpace(pair[1]))
		}
	}

	for _, a := range flag.Args() {
		if !strings.Contains(a, "=") {
			fmt.Println("Context attribute \"" + a + "\" does not contain \"=\".")
			os.Exit(1)
		}
		pair := strings.SplitN(a, "=", 2)
		value, err := dfl.Parse(strings.TrimSpace(pair[1]))
		if err != nil {
			fmt.Println(errors.Wrap(err, "Could not parse context variable"))
			os.Exit(1)
		}
		value = value.Compile()
		switch value.(type) {
		case dfl.Literal:
			ctx[strings.TrimSpace(pair[0])] = value.(dfl.Literal).Value
		default:
			ctx[strings.TrimSpace(pair[0])] = dfl.TryConvertString(pair[1])
		}
	}

	root, err := dfl.Parse(filter_text)
	if err != nil {
		fmt.Println("Error parsing filter expression.")
		fmt.Println(err)
		os.Exit(1)
	}

	if verbose {
		fmt.Println("******************* Context *******************")
		out, err := yaml.Marshal(ctx)
		if err != nil {
			fmt.Println("Error marshaling context to yaml.")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}

	if verbose {
		fmt.Println("******************* Parsed *******************")
		out, err := yaml.Marshal(root.Map())
		if err != nil {
			fmt.Println("Error marshaling expression to yaml.")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}

	root = root.Compile()

	if verbose {
		fmt.Println("******************* Compiled *******************")
		out, err := yaml.Marshal(root.Map())
		if err != nil {
			fmt.Println("Error marshaling expression to yaml.")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(out))
	}

	result, err := root.Evaluate(ctx, dfl.NewFuntionMapWithDefaults())
	if err != nil {
		fmt.Println("Error evaluating expression.")
		fmt.Println(err)
		os.Exit(1)
	}

	switch result.(type) {
	case bool:
		result_bool := result.(bool)
		if verbose {
			fmt.Println("******************* Result *******************")
			if result_bool {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}

			elapsed := time.Since(start)
			fmt.Println("Done in " + elapsed.String())
		}

		if result_bool {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	default:
		os.Exit(1)
	}
}
