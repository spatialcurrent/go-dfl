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
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-stringify/pkg/stringify"

	"github.com/spatialcurrent/go-dfl/pkg/cli"
	"github.com/spatialcurrent/go-dfl/pkg/dfl"
	"github.com/spatialcurrent/go-dfl/pkg/shell"
)

//"github.com/atotto/clipboard"

//"github.com/ahmetb/go-cursor"

var DefaultQuotes = []string{"'", "\"", "`"}

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

			inputReader, _, err := grw.ReadFromResource(inputUri, "", 4096, nil)
			if err != nil {
				return errors.Wrapf(err, "error reading dfl file at uri %q", inputUri)
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
	cli.InitFmtFlags(fmtCommand.Flags())
	rootCommand.AddCommand(fmtCommand)

	execCommand := &cobra.Command{
		Use:   "exec",
		Short: "executes a DFL expression (parse, compile, and evalutes)",
		Long:  "executes a DFL expression (parse, compile, and evalutes)",
		RunE: func(cmd *cobra.Command, args []string) error {

			funcs := dfl.DefaultFunctionMap
			quotes := dfl.DefaultQuotes

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

			inputUri := v.GetString("input-uri")

			inputReader, _, err := grw.ReadFromResource(inputUri, "", 4096, nil)
			if err != nil {
				return errors.Wrapf(err, "error reading dfl file at uri %q", inputUri)
			}

			inputBytes, err := inputReader.ReadAllAndClose()
			if err != nil {
				return err
			}

			node, err := dfl.ParseCompile(strings.TrimSpace(dfl.RemoveComments(string(inputBytes))))
			if err != nil {
				return err
			}

			ctx := map[string]interface{}{}

			if v.GetBool("env") {
				for _, e := range os.Environ() {
					pair := strings.SplitN(e, "=", 2)
					ctx[strings.TrimSpace(pair[0])] = dfl.TryConvertString(strings.TrimSpace(pair[1]))
				}
			}

			if v.GetBool("args") {
				m, errArgs := cli.ParseContextArguments(args, funcs, quotes)
				if errArgs != nil {
					return errors.Wrap(errArgs, "error parsing context from arguments")
				}
				for k, v := range m {
					ctx[k] = v
				}
			}

			vars := map[string]interface{}{}
			if str := strings.TrimSpace(v.GetString("vars")); len(str) > 0 {
				m, errVars := cli.ParseInitialVariables(str, funcs, quotes)
				if errVars != nil {
					return errors.Wrap(errVars, "error parsing initial variables")
				}
				vars = m
			}

			_, result, err := node.Evaluate(vars, ctx, funcs, quotes)
			if err != nil {
				return errors.Wrap(err, "error evaluating")
			}

			result, err = stringify.StringifyMapKeys(result, stringify.NewDefaultStringer())
			if err != nil {
				return errors.Wrap(err, "error stringifying result")
			}

			outBytes, err := shell.FormatOutput(v, result, quotes)
			if err != nil {
				return errors.Wrap(err, "error formatting output")
			}

			outputUri := v.GetString("output-uri")

			outputWriter, err := grw.WriteToResource(outputUri, "none", v.GetBool("append"), nil)
			if err != nil {
				return errors.Wrapf(err, "error writing dfl file to uri %q", outputUri)
			}

			_, err = outputWriter.Write(outBytes)
			if err != nil {
				return errors.Wrapf(err, "error writing dfl file to uri %q", outputUri)
			}

			err = outputWriter.Flush()
			if err != nil {
				return errors.Wrapf(err, "error flushing dfl file to uri %q", outputUri)
			}

			err = outputWriter.Close()
			if err != nil {
				return errors.Wrapf(err, "error writing dfl file to uri %q", outputUri)
			}

			return nil
		},
	}
	flags := execCommand.Flags()
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

	shellCommand := &cobra.Command{
		Use:   "shell",
		Short: "executes a DFL expression (parse, compile, and evalutes)",
		Long:  "executes a DFL expression (parse, compile, and evalutes)",
		RunE: func(cmd *cobra.Command, args []string) error {

			quotes := dfl.DefaultQuotes

			v := viper.New()
			err := v.BindPFlags(cmd.Flags())
			if err != nil {
				return err
			}
			v.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
			v.AutomaticEnv()

			vars := map[string]interface{}{}

			if v.GetBool("env") {
				for _, e := range os.Environ() {
					pair := strings.SplitN(e, "=", 2)
					vars[strings.TrimSpace(pair[0])] = dfl.TryConvertString(strings.TrimSpace(pair[1]))
				}
			}

			if v.GetBool("args") {
				m, errArgs := cli.ParseContextArguments(args, dfl.DefaultFunctionMap, quotes)
				if errArgs != nil {
					return errors.Wrap(errArgs, "error parsing context from arguments")
				}
				for k, v := range m {
					vars[k] = v
				}
			}

			errKeyboard := keyboard.Open()
			if errKeyboard != nil {
				return errors.Wrap(errKeyboard, "error creating keyboard")
			}
			defer keyboard.Close()

			s := shell.New()

			s.Executor.Funcs["read"] = func(funcs dfl.FunctionMap, vars map[string]interface{}, ctx interface{}, args []interface{}, quotes []string) (interface{}, error) {
				fmt.Println("Args:", args)
				if len(args) == 0 {
					return make([]byte, 0), nil
				}
				path, err := homedir.Expand(fmt.Sprint(args[0]))
				if err != nil {
					return make([]byte, 0), err
				}
				fmt.Println("Path:", path)
				return ioutil.ReadFile(path)
			}

			s.PrintHeader()

			for {

				s.ResetLine()
				s.UpdateScreen()

				for {

					char, key, err := s.GetKey()
					if err != nil {
						return errors.Wrap(err, "error reading input from keyboard")
					}

					if key == keyboard.KeyCtrlC {
						fmt.Print("^C\n")
						return nil
					}

					if key == keyboard.KeyCtrlZ {
						fmt.Print("^Z\n")
						return nil
					}

					if key == keyboard.KeyCtrlX {
						err := s.CutLine()
						if err != nil {
							return errors.Wrap(err, "error writing to clipboard")
						}
						continue
					}

					if key == keyboard.KeyCtrlV {
						err := s.PasteLine()
						if err != nil {
							return errors.Wrap(err, "error reading from clipboard")
						}
						continue
					}

					if key == keyboard.KeyEnter {
						fmt.Print("\n")
						lineTrimmed := strings.TrimSpace(s.Line)
						if len(lineTrimmed) > 0 && lineTrimmed[len(lineTrimmed)-1] == '|' {
							//s.Line = lineTrimmed + " "
							s.Line = lineTrimmed + "\n"
							continue
						}
						break
					}

					if key == keyboard.KeySpace {
						s.WriteString(" ")
						continue
					}

					if key == keyboard.KeyArrowLeft {
						s.MoveLeft()
						continue
					}

					if key == keyboard.KeyArrowRight {
						s.MoveRight()
						continue
					}

					if key == keyboard.KeyArrowDown {
						if s.History.Cursor == 0 {
							continue
						}
						if s.History.Cursor == 1 {
							s.History.Forward()
							s.SetLine("")
						} else if s.History.Cursor > 1 {
							s.History.Forward()
							s.SetLine(s.History.Line())
						}
						s.ClearLine()
						s.UpdateScreen()
					}

					if key == keyboard.KeyArrowUp {
						if s.History.Cursor < s.History.Len() {
							s.History.Back()
							s.SetLine(s.History.Line())
							s.ClearLine()
							s.UpdateScreen()
						}
						continue
					}

					switch key {
					case keyboard.KeyDelete, keyboard.KeyBackspace, keyboard.KeyBackspace2, keyboard.KeyEsc:
						switch key {
						case keyboard.KeyDelete, keyboard.KeyBackspace, keyboard.KeyBackspace2:
							if len(s.Line) > 0 {
								s.History.Cursor = 0
								s.Backspace()
							}
						case keyboard.KeyEsc:
							s.History.Cursor = 0
							s.Line = ""
							s.ClearLine()
							s.UpdateScreen()
						}
					default:
						s.WriteString(string(char))
					}

				}

				lineTrimmed := s.CleanLine()

				if len(lineTrimmed) > 0 {

					if lineTrimmed == "exit" || lineTrimmed == "quit" {
						return nil
					}

					if lineTrimmed == "help" {
						s.PrintHelp()
						continue
					}

					if lineTrimmed == "history" || lineTrimmed == "!" {
						s.PrintHistory()
						continue
					}

					if lineTrimmed == "examples" {
						s.PrintExamples()
						continue
					}

					if lineTrimmed == "funcs" || lineTrimmed == "functions" {
						s.PrintFunctions()
						continue
					}

					if lineTrimmed == "!!" {
						lineTrimmed = s.History.Last()
						fmt.Println(lineTrimmed)
					} else if lineTrimmed[0] == '!' {
						cursor, err := strconv.Atoi(lineTrimmed[1:])
						if err != nil {
							fmt.Fprintln(os.Stderr, errors.Wrapf(err, "command %q not found", lineTrimmed[1:]))
							continue
						}
						if cursor > 0 && cursor <= s.History.Len() {
							lineTrimmed = s.History.Get(cursor - 1)
							fmt.Println(lineTrimmed)
						} else {
							fmt.Fprintln(os.Stderr, errors.Wrapf(err, "command %d not found", cursor))
							continue
						}
					} else {
						if s.History.Empty() || s.History.Last() != lineTrimmed {
							s.History.Push(lineTrimmed)
						}
					}

					outputObject, err := s.Exec(lineTrimmed)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						continue
					}

					if outputObject == nil {
						continue
					}

					outputBytes, err := shell.FormatOutput(v, outputObject, quotes)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						continue
					}

					fmt.Println(string(outputBytes))

				}

			}
		},
	}
	flags = shellCommand.Flags()
	cli.InitOutputFlags(flags)
	rootCommand.AddCommand(shellCommand)

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}

}
