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
	"os"
	"strconv"
	"strings"
)

import (
	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

import (
	"github.com/spatialcurrent/go-dfl/pkg/cli"
	"github.com/spatialcurrent/go-reader-writer/grw"
	"github.com/spatialcurrent/go-stringify/pkg/stringify"
)

import (
	"github.com/spatialcurrent/go-dfl/pkg/dfl"
	"github.com/spatialcurrent/go-dfl/pkg/shell"
)

var DefaultQuotes = []string{"'", "\"", "`"}

func main() {

	funcs := dfl.DefaultFunctionMap
	quotes := dfl.DefaultQuotes

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
	cli.InitFmtFlags(fmtCommand.Flags())
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

			inputUri := v.GetString("input-uri")

			inputReader, _, err := grw.ReadFromResource(inputUri, "", 4096, false, nil)
			if err != nil {
				return errors.Wrap(err, "error reading dfl file at uri: "+inputUri)
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
				return errors.Wrap(err, "error writing dfl file to uri: "+outputUri)
			}

			_, err = outputWriter.Write(outBytes)
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
				m, errArgs := cli.ParseContextArguments(args, funcs, quotes)
				if errArgs != nil {
					return errors.Wrap(errArgs, "error parsing context from arguments")
				}
				for k, v := range m {
					vars[k] = v
				}
			}

			history := make([]string, 0)

			errKeyboard := keyboard.Open()
			if errKeyboard != nil {
				return errors.Wrap(errKeyboard, "error creating keyboard")
			}
			defer keyboard.Close()

			for {

				err := shell.UpdateLine(v, vars, quotes, "")
				if err != nil {
					fmt.Fprintln(os.Stderr, err.Error())
				}

				cursor := 0
				line := ""
				for {
					char, key, err := keyboard.GetKey()

					if err != nil {
						return errors.Wrap(err, "error reading input from keyboard")
					}

					if key == keyboard.KeyCtrlC {
						fmt.Print("\n")
						return nil
					}

					if key == keyboard.KeyCtrlX {
						err := clipboard.WriteAll(line)
						if err != nil {
							return errors.Wrap(err, "error writing to clipboard")
						}
						line = ""
						shell.ClearLine()
						err = shell.UpdateLine(v, vars, quotes, "")
						if err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
						}
						continue
					}

					if key == keyboard.KeyCtrlV {
						p, err := clipboard.ReadAll()
						if err != nil {
							return errors.Wrap(err, "error reading from clipboard")
						}
						fmt.Print(p)
						line += p
						continue
					}

					if key == keyboard.KeyEnter {
						fmt.Print("\n")
						lineTrimmed := strings.TrimSpace(line)
						if len(lineTrimmed) > 0 && lineTrimmed[len(lineTrimmed)-1] == '|' {
							line = lineTrimmed + " "
							continue
						}
						break
					}

					if key == keyboard.KeySpace {
						line += " "
						fmt.Print(" ")
						continue
					}

					switch key {
					case keyboard.KeyArrowUp, keyboard.KeyArrowDown, keyboard.KeyArrowLeft, keyboard.KeyArrowRight, keyboard.KeyDelete, keyboard.KeyBackspace, keyboard.KeyBackspace2, keyboard.KeyEsc:
						switch key {
						case keyboard.KeyArrowLeft, keyboard.KeyArrowRight:
						case keyboard.KeyArrowUp:
							if cursor < len(history) {
								cursor += 1
								line = history[len(history)-cursor]
							}
						case keyboard.KeyArrowDown:
							if cursor == 1 {
								cursor--
								line = ""
							} else if cursor > 1 {
								cursor--
								line = history[len(history)-cursor]
							}
						case keyboard.KeyDelete, keyboard.KeyBackspace, keyboard.KeyBackspace2:
							if len(line) > 0 {
								cursor = 0
								line = line[0 : len(line)-1]
							}
						case keyboard.KeyEsc:
							cursor = 0
							line = ""
						}
						shell.ClearLine()
						err := shell.UpdateLine(v, vars, quotes, line)
						if err != nil {
							fmt.Fprintln(os.Stderr, err.Error())
						}
					default:
						fmt.Print(string(char))
						line += string(char)
					}

				}

				lineTrimmed := strings.TrimSpace(dfl.RemoveComments(string(line)))

				//fmt.Println("line:", line)
				//fmt.Println("line:", []byte(line))
				if len(lineTrimmed) > 0 {

					if lineTrimmed == "exit" || lineTrimmed == "quit" {
						return nil
					}

					if lineTrimmed == "help" {
						shell.PrintHelp()
						continue
					}

					if lineTrimmed == "history" || lineTrimmed == "!" {
						shell.PrintHistory(history)
						continue
					}

					if lineTrimmed == "examples" {
						shell.PrintExamples()
						continue
					}

					if lineTrimmed == "funcs" || lineTrimmed == "functions" {
						shell.PrintFunctions(funcs)
						continue
					}

					if lineTrimmed[0] == '!' {
						cursor, err := strconv.Atoi(lineTrimmed[1:])
						if err != nil {
							fmt.Fprintln(os.Stderr, errors.Wrapf(err, "command %q not found", lineTrimmed[1:]))
							continue
						}
						if cursor > 0 && cursor <= len(history) {
							lineTrimmed = history[cursor-1]
						} else {
							fmt.Fprintln(os.Stderr, errors.Wrapf(err, "command %d not found", cursor))
							continue
						}
					} else {
						if len(history) == 0 || history[len(history)-1] != lineTrimmed {
							history = append(history, lineTrimmed)
						}
					}

					node, err := dfl.ParseCompile(lineTrimmed)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						continue
					}

					newVars, ctx, err := node.Evaluate(vars, nil, funcs, quotes)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						continue
					}

					vars = newVars

					if ctx == nil {
						continue
					}

					outputBytes, err := shell.FormatOutput(v, ctx, quotes)
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						continue
					}

					fmt.Println(string(outputBytes))

				}

			}
			return nil

		},
	}
	flags = shellCommand.Flags()
	cli.InitOutputFlags(flags)
	rootCommand.AddCommand(shellCommand)

	if err := rootCommand.Execute(); err != nil {
		panic(err)
	}

}
