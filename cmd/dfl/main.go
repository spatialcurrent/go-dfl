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
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

import (
	"github.com/spatialcurrent/go-dfl/dfl"
	"github.com/spatialcurrent/go-reader-writer/grw"
	stringify "github.com/spatialcurrent/go-stringify"
)

var DefaultQuotes = []string{"'", "\"", "`"}

func parseContextArguments(args []string, funcs dfl.FunctionMap) (map[string]interface{}, error) {
	ctx := map[string]interface{}{}
	for _, a := range args {
		if !strings.Contains(a, "=") {
			return ctx, errors.New("Context attribute \"" + a + "\" does not contain \"=\".")
		}
		pair := strings.SplitN(a, "=", 2)
		value, _, err := dfl.Parse(strings.TrimSpace(pair[1]))
		if err != nil {
			return ctx, errors.Wrap(err, "Could not parse context variable")
		}
		value = value.Compile()
		switch value.(type) {
		case dfl.Array:
			_, arr, err := value.(dfl.Array).Evaluate(map[string]interface{}{}, map[string]interface{}{}, funcs, DefaultQuotes[1:])
			if err != nil {
				return ctx, errors.Wrap(err, "error evaluating context expression for "+strings.TrimSpace(pair[0]))
			}
			ctx[strings.TrimSpace(pair[0])] = arr
		case dfl.Set:
			_, arr, err := value.(dfl.Set).Evaluate(map[string]interface{}{}, map[string]interface{}{}, funcs, DefaultQuotes[1:])
			if err != nil {
				return ctx, errors.Wrap(err, "error evaluating context expression for "+strings.TrimSpace(pair[0]))
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
	return ctx, nil
}

func parseInitialVariables(str string, funcs dfl.FunctionMap, quotes []string) (map[string]interface{}, error) {
	vars := map[string]interface{}{}
	if len(str) > 0 {
		_, result, err := dfl.ParseCompileEvaluateMap(
			str,
			dfl.NoVars,
			dfl.NoContext,
			funcs,
			quotes)
		if err != nil {
			return vars, errors.Wrap(err, "error parsing initial dfl vars as map")
		}
		if m, ok := stringify.StringifyMapKeys(result).(map[string]interface{}); ok {
			vars = m
		}
	}
	return vars, nil
}

func main() {

	rootCommand := cobra.Command{
		Use:   "dfl [flags]",
		Short: "CLI for Dynamic Filter Language",
		Long:  "CLI for Dynamic Filter Language",
	}

	completionCommandLong := ""
	if _, err := os.Stat("/etc/bash_completion.d/"); !os.IsNotExist(err) {
		completionCommandLong = "To install completion scripts run:\ndfl completion > /etc/bash_completion.d/dfl"
	} else {
		if _, err := os.Stat("/usr/local/etc/bash_completion.d/"); !os.IsNotExist(err) {
			completionCommandLong = "To install completion scripts run:\ndfl completion > /usr/local/etc/bash_completion.d/dfl"
		} else {
			completionCommandLong = "To install completion scripts run:\ndfl completion > .../bash_completion.d/dfl"
		}
	}

	completionCommand := &cobra.Command{
		Use:   "completion",
		Short: "Generates bash completion scripts",
		Long:  completionCommandLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return rootCommand.GenBashCompletion(os.Stdout)
		},
	}
	rootCommand.AddCommand(completionCommand)

	fmtCommand := &cobra.Command{
		Use:   "fmt",
		Short: "Formats a dfl expression",
		Long:  "Formats a dfl expression",
		RunE: func(cmd *cobra.Command, args []string) error {

			err := cmd.ParseFlags(args)
			if err != nil {
				return err
			}

			flag := cmd.Flags()

			v := viper.New()
			err = v.BindPFlags(flag)
			if err != nil {
				return err
			}
			v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			v.AutomaticEnv()

			inputUri := v.GetString("uri")

			inputReader, _, err := grw.ReadFromResource(inputUri, "", 4096, false, nil)
			if err != nil {
				return errors.Wrap(err, "error reading dfl file at uri: "+inputUri)
			}

			inputBytes, err := inputReader.ReadAllAndClose()
			if err != nil {
				return err
			}

			node, _, err := dfl.Parse(strings.TrimSpace(dfl.RemoveComments(string(inputBytes))))
			if err != nil {
				return err
			}

			if v.GetBool("compile") {
				node = node.Compile()
			}

			out := node.Dfl(dfl.DefaultQuotes, v.GetBool("pretty"), v.GetInt("tabs"))

			fmt.Println(out)

			return nil
		},
	}
	flags := fmtCommand.Flags()
	flags.StringP("uri", "u", "stdin", "uri to DFL file")
	flags.BoolP("compile", "c", false, "compile expression")
	flags.BoolP("pretty", "p", false, "pretty output")
	flags.IntP("tabs", "t", 0, "tabs")
	rootCommand.AddCommand(fmtCommand)

	execCommand := &cobra.Command{
		Use:   "exec",
		Short: "executes a DFL expression (parse, compile, and evalutes)",
		Long:  "executes a DFL expression (parse, compile, and evalutes)",
		RunE: func(cmd *cobra.Command, args []string) error {

			err := cmd.ParseFlags(args)
			if err != nil {
				return err
			}

			flag := cmd.Flags()

			v := viper.New()
			err = v.BindPFlags(flag)
			if err != nil {
				return err
			}
			v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			v.AutomaticEnv()

			inputUri := v.GetString("uri")

			inputReader, _, err := grw.ReadFromResource(inputUri, "", 4096, false, nil)
			if err != nil {
				return errors.Wrap(err, "error reading dfl file at uri: "+inputUri)
			}

			inputBytes, err := inputReader.ReadAllAndClose()
			if err != nil {
				return err
			}

			node, _, err := dfl.Parse(strings.TrimSpace(dfl.RemoveComments(string(inputBytes))))
			if err != nil {
				return err
			}

			node = node.Compile()

			funcs := dfl.DefaultFunctionMap
			quotes := dfl.DefaultQuotes

			ctx := map[string]interface{}{}

			if v.GetBool("env") {
				for _, e := range os.Environ() {
					pair := strings.SplitN(e, "=", 2)
					ctx[strings.TrimSpace(pair[0])] = dfl.TryConvertString(strings.TrimSpace(pair[1]))
				}
			}

			if v.GetBool("args") {
				m, errArgs := parseContextArguments(args, funcs)
				if errArgs != nil {
					return errors.Wrap(errArgs, "error parsing context from arguments")
				}
				for k, v := range m {
					ctx[k] = v
				}
			}

			vars := map[string]interface{}{}
			if str := strings.TrimSpace(v.GetString("vars")); len(str) > 0 {
				m, errVars := parseInitialVariables(str, funcs, quotes)
				if errVars != nil {
					return errors.Wrap(errVars, "error parsing initial variables")
				}
				vars = m
			}

			_, result, err := node.Evaluate(vars, ctx, funcs, quotes)
			if err != nil {
				return errors.Wrap(err, "error evaluating")
			}

			result = stringify.StringifyMapKeys(result)

			pretty := v.GetBool("pretty")

			outString := ""
			if v.GetBool("json") {
				if pretty {
					outputBytes, errJSON := json.MarshalIndent(result, "", "  ")
					if errJSON != nil {
						return errors.Wrap(errJSON, "error marshalling result")
					}
					outString = string(outputBytes)
				} else {
					outputBytes, errJSON := json.Marshal(result)
					if errJSON != nil {
						return errors.Wrap(errJSON, "error marshalling result")
					}
					outString = string(outputBytes)
				}
			} else if v.GetBool("yaml") {
				outputBytes, errYAML := yaml.Marshal(result)
				if errYAML != nil {
					return errors.Wrap(errYAML, "error marshalling result")
				}
				outString = string(outputBytes)
			} else {
				outString = dfl.TryFormatLiteral(result, quotes, v.GetBool("pretty"), v.GetInt("tabs"))
			}

			outputUri := v.GetString("output-uri")

			outputWriter, err := grw.WriteToResource(outputUri, "none", v.GetBool("append"), nil)
			if err != nil {
				return errors.Wrap(err, "error writing dfl file to uri: "+outputUri)
			}

			_, err = outputWriter.WriteString(outString)
			if err != nil {
				return errors.Wrap(err, "error writing dfl file to uri: "+outputUri)
			}

			err = outputWriter.Close()
			if err != nil {
				return errors.Wrap(err, "error writing dfl file to uri: "+outputUri)
			}

			return nil
		},
	}
	flags = execCommand.Flags()
	flags.StringP("input-uri", "i", "stdin", "input uri to DFL file")
	flags.StringP("output-uri", "o", "stdout", "output uri to DFL file")
	flags.StringP("vars", "v", "", "vars")
	flags.BoolP("args", "a", false, "load context attributes from arguments")
	flags.BoolP("env", "e", false, "load context attributes from environment variables")
	flags.BoolP("pretty", "p", false, "pretty output")
	flags.BoolP("json", "j", false, "json output")
	flags.BoolP("yaml", "y", false, "yaml output")
	flags.Bool("append", false, "append output")
	flags.IntP("tabs", "t", 0, "tabs")
	rootCommand.AddCommand(execCommand)

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}

}
