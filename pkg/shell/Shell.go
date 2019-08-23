package shell

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"

	"github.com/spatialcurrent/go-dfl/pkg/dfl"
)

const (
	moveLeft       = "\033[D" // https://www.student.cs.uwaterloo.ca/~cs452/terminal.html
	moveRight      = "\033[C" // https://www.student.cs.uwaterloo.ca/~cs452/terminal.html
	saveCursor     = "\033[s" // http://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html
	restoreCursor  = "\033[u" // http://tldp.org/HOWTO/Bash-Prompt-HOWTO/x361.html
	clearLineRight = "\033[K" // https://www.student.cs.uwaterloo.ca/~cs452/terminal.html
)

type Executor struct {
	Vars  map[string]interface{}
	Funcs dfl.FunctionMap
}

type Formatter struct {
	Quotes []string
}

type Shell struct {
	History   *History
	Executor  *Executor
	Formatter *Formatter
	Line      string
	Column    int
}

func (s *Shell) PrintHeader() {
	fmt.Println("DFL Shell")
	fmt.Println("Type \"help\" for more information.")
	//Type "help", "copyright", "credits" or "license" for more information.
}

func (s *Shell) PrintHelp() {
	fmt.Println("# Help")
	fmt.Println("# examples: print examples")
	fmt.Println("# history: print history")
	fmt.Println("# funcs, functions: print function names")
	fmt.Println("# Ctrl-C, exit, quit: exit shell")
	fmt.Println("# Ctrl-X: cut current line")
	fmt.Println("# Ctrl-V: paste from clipboard")
}

func (s *Shell) PrintHistory() {
	fmt.Println("# History")
	for i, line := range s.History.Lines {
		fmt.Println(strconv.Itoa(i+1) + "\t" + line)
	}
}

func (s *Shell) PrintFunctions() {
	fmt.Println("# Functions")
	lines := make([]string, 0, len(s.Executor.Funcs))
	for k, _ := range s.Executor.Funcs {
		lines = append(lines, k)
	}
	sort.Strings(lines)
	fmt.Println(strings.Join(lines, "\n"))
}

func (s *Shell) PrintExamples() {
	fmt.Println("# Examples")
	fmt.Println(`
# Add Numbers
$x := 10
$y := 20
$x + $y`)
	fmt.Println("")
}

func (s *Shell) MoveLeft() {
	if s.Column > 0 {
		s.Column--
		fmt.Print(moveLeft)
	}
}

func (s *Shell) MoveRight() {
	if s.Column < len(s.Line) {
		s.Column += 1
		fmt.Print(moveRight)
	}
}

func (s *Shell) GetKey() (rune, keyboard.Key, error) {
	return keyboard.GetKey()
}

func (s *Shell) ClearLine() {
	fmt.Print("\r")           // carriage return to beginning of line
	fmt.Print(clearLineRight) // https://www.student.cs.uwaterloo.ca/~cs452/terminal.html
}

func (s *Shell) UpdateScreen() {
	vars := dfl.TryFormatLiteral(s.Executor.Vars, s.Formatter.Quotes, false, 0)
	_, _ = fmt.Fprintf(os.Stdout, "(%s) > %s", vars, s.Line)
}

func (s *Shell) SetLine(str string) {
	s.Line = str
	s.Column = len(str)
}

func (s *Shell) CutLine() error {
	err := clipboard.WriteAll(s.Line)
	if err != nil {
		return err
	}
	s.ResetLine()
	s.ClearLine()
	s.UpdateScreen()
	return nil
}

func (s *Shell) PasteLine() error {
	p, err := clipboard.ReadAll()
	if err != nil {
		return err
	}
	s.WriteString(p)
	return nil
}

func (s *Shell) Exec(line string) (interface{}, error) {

	node, err := dfl.ParseCompile(line)
	if err != nil {
		return nil, err
	}

	newVars, out, err := node.Evaluate(s.Executor.Vars, nil, s.Executor.Funcs, s.Formatter.Quotes)
	if err != nil {
		return nil, err
	}

	s.Executor.Vars = newVars

	return out, err
}

func (s *Shell) ResetLine() {
	s.Line = ""
	s.Column = 0
}

func (s *Shell) WriteString(str string) {
	if s.Column == len(s.Line) {
		s.Line += str
		s.Column += len(str)
		fmt.Print(str)
	} else {
		if s.Column == 0 {
			s.Line = str + s.Line
			s.Column += len(str)
			fmt.Print(str)
		} else {
			s.Line = s.Line[0:s.Column] + str + s.Line[s.Column:]
			s.Column += len(str)
			fmt.Print(str)
		}
	}
}

func (s *Shell) Backspace() {
	if len(s.Line) > 0 {
		if s.Column > 0 {

			fmt.Print(moveLeft)
			fmt.Print(saveCursor)
			fmt.Print(clearLineRight) // https://www.student.cs.uwaterloo.ca/~cs452/terminal.html

			if s.Column == 1 {
				s.Line = s.Line[1:]
			} else {
				if s.Column == len(s.Line) {
					s.Line = s.Line[0 : len(s.Line)-1]
				} else {
					s.Line = s.Line[0:s.Column-1] + s.Line[s.Column:]
				}
			}

			s.Column--
			fmt.Print(s.Line[s.Column:])
			fmt.Print(restoreCursor)
		}
	}
}

func (s *Shell) CleanLine() string {
	return strings.TrimSpace(dfl.RemoveComments(string(s.Line)))
}

func New() *Shell {
	return &Shell{
		History: &History{
			Lines:  make([]string, 0),
			Cursor: 0,
		},
		Executor: &Executor{
			Vars:  map[string]interface{}{},
			Funcs: dfl.DefaultFunctionMap,
		},
		Formatter: &Formatter{
			Quotes: dfl.DefaultQuotes,
		},
		Line: "",
	}
}
