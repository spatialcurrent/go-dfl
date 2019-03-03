// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

// dfl is the command line program for DFL.
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
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
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

var GO_DFL_DEFAULT_QUOTES = []string{"'", "\"", "`"}

func main() {

	start := time.Now()

	var filter_text string

	var load_env bool
	var verbose bool
	var version bool
	var sql bool
	var pretty bool
	var dry_run bool
	var help bool

	flag.StringVar(&filter_text, "f", "", "The DFL expression to evaulate")
	flag.BoolVar(&sql, "sql", false, "Prints SQL version of expression to stdout")
	flag.BoolVar(&pretty, "pretty", false, "Prints pretty version of expression to stdout")

	flag.BoolVar(&load_env, "env", false, "Load environment variables")
	flag.BoolVar(&verbose, "verbose", false, "Provide verbose output")
	flag.BoolVar(&version, "version", false, "Prints version to stdout")
	flag.BoolVar(&dry_run, "dry_run", false, "Do a dry run (parse and compile expression but do not evaluate)")
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
		fmt.Println(dfl.Version)
		os.Exit(0)
	}

	ctx := map[string]interface{}{}

	if load_env {
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			ctx[strings.TrimSpace(pair[0])] = dfl.TryConvertString(strings.TrimSpace(pair[1]))
		}
	}

	funcs := dfl.NewFuntionMapWithDefaults()

	for _, a := range flag.Args() {
		if !strings.Contains(a, "=") {
			log.Fatal(errors.New("Context attribute \"" + a + "\" does not contain \"=\"."))
		}
		pair := strings.SplitN(a, "=", 2)
		value, _, err := dfl.Parse(strings.TrimSpace(pair[1]))
		if err != nil {
			log.Fatal(errors.Wrap(err, "Could not parse context variable"))
		}
		value = value.Compile()
		switch value.(type) {
		case dfl.Array:
			_, arr, err := value.(dfl.Array).Evaluate(map[string]interface{}{}, map[string]interface{}{}, funcs, GO_DFL_DEFAULT_QUOTES[1:])
			if err != nil {
				log.Fatal(errors.Wrap(err, "error evaluating context expression for "+strings.TrimSpace(pair[0])))
			}
			ctx[strings.TrimSpace(pair[0])] = arr
		case dfl.Set:
			_, arr, err := value.(dfl.Set).Evaluate(map[string]interface{}{}, map[string]interface{}{}, funcs, GO_DFL_DEFAULT_QUOTES[1:])
			if err != nil {
				log.Fatal(errors.Wrap(err, "error evaluating context expression for "+strings.TrimSpace(pair[0])))
			}
			ctx[strings.TrimSpace(pair[0])] = arr
		case dfl.Literal:
			ctx[strings.TrimSpace(pair[0])] = value.(dfl.Literal).Value
		case *dfl.Literal:
			ctx[strings.TrimSpace(pair[0])] = value.(*dfl.Literal).Value
		default:
			ctx[strings.TrimSpace(pair[0])] = dfl.TryConvertString(pair[1])
		}
	}

	root, _, err := dfl.Parse(filter_text)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error parsing filter expression"))
	}

	if pretty {
		fmt.Println("# Pretty Version \n" + root.Dfl(dfl.DefaultQuotes[1:], true, 0) + "\n")
	}

	if sql {
		fmt.Println("# SQL Version \n" + root.Sql(pretty, 0) + "\n")
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

		fmt.Println("******************* Parsed *******************")
		out, err = yaml.Marshal(root.Map())
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
		fmt.Println("# YAML Version\n" + string(out))
		fmt.Println("# DFL Version\n" + GO_DFL_DEFAULT_QUOTES[0] + root.Dfl(GO_DFL_DEFAULT_QUOTES[1:], false, 0) + GO_DFL_DEFAULT_QUOTES[0])
		if sql {
			fmt.Println("# SQL Version\n" + root.Sql(pretty, 0) + "\n")
		}
	}

	if dry_run {
		os.Exit(0)
	}

	_, result, err := root.Evaluate(map[string]interface{}{}, ctx, funcs, GO_DFL_DEFAULT_QUOTES[1:])
	if err != nil {
		log.Fatal(errors.Wrap(err, "error evaluating expression"))
	}

	switch result.(type) {
	case bool:
		result_bool := result.(bool)
		if verbose {
			fmt.Println("******************* Result *******************")
			fmt.Println(dfl.TryFormatLiteral(result, GO_DFL_DEFAULT_QUOTES[1:], false, 0))
			elapsed := time.Since(start)
			fmt.Println("Done in " + elapsed.String())
		}
		if result_bool {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	default:
		if verbose {
			fmt.Println("******************* Result *******************")
			fmt.Println("Type:", reflect.TypeOf(result))
			fmt.Println("Value:", dfl.TryFormatLiteral(result, GO_DFL_DEFAULT_QUOTES[1:], false, 0))
			elapsed := time.Since(start)
			fmt.Println("Done in " + elapsed.String())
		}
	}
}
