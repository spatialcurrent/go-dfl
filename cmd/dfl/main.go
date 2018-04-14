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
)

import (
	"github.com/spatialcurrent/go-dfl/dfl"
)

var GO_DFL_VERSION = "0.0.1"

func dfl_build_funcs() map[string]func(map[string]interface{}, []string) (interface{}, error) {
	funcs := map[string]func(map[string]interface{}, []string) (interface{}, error){}

	funcs["len"] = func(ctx map[string]interface{}, args []string) (interface{}, error) {
		if len(args) != 1 {
			return 0, errors.New("Invalid number of arguments to len.")
		}
		return len(args[0]), nil
	}

	return funcs
}

func main() {

	start := time.Now()

	var filter_text string

	var verbose bool
	var version bool
	var help bool

	flag.StringVar(&filter_text, "filter", "", "The DFL expression to evaulate")

	flag.BoolVar(&verbose, "verbose", false, "Provide verbose output")
	flag.BoolVar(&version, "version", false, "Prints version to stdout")
	flag.BoolVar(&help, "help", false, "Print help")

	flag.Parse()

	if help {
		fmt.Println("Usage: dfl -filter INPUT [-verbose] [-version] [-help] [A=1] [B=2]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: Provided no arguments.")
		fmt.Println("Run \"dfl -help\" for more information.")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Println("Usage: dfl -filter INPUT [-verbose] [-version] [-help] [A=1] [B=2]")
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println(GO_DFL_VERSION)
		os.Exit(0)
	}

	ctx := map[string]interface{}{}
	for _, a := range flag.Args() {
		parts := strings.SplitN(a, "=", 2)
		ctx[parts[0]] = dfl.TryConvertString(parts[1])
	}

	root, err := dfl.Parse(filter_text)
	if err != nil {
		fmt.Println("Error parsing filter expression.")
		fmt.Println(err)
		os.Exit(1)
	}

	funcs := dfl_build_funcs()
	result, err := root.Evaluate(ctx, funcs)
	if err != nil {
		fmt.Println("Error evaluating expression.")
		fmt.Println(err)
		os.Exit(1)
	}

	switch result.(type) {
	case bool:
		result_bool := result.(bool)
		if verbose {
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
