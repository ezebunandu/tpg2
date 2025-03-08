package shell

import (
	"errors"
	"io"
	"os/exec"
	"strings"
    "fmt"
    "bufio"
)

func CmdFromString(input string) (*exec.Cmd, error) {
    args := strings.Fields(input)
    if len(args) < 1 {
        return nil, errors.New("empty input")
    }
    return exec.Command(args[0], args[1:]...), nil
}

type Session struct {
    Stdin io.Reader
    Stdout io.Writer
    Stderr io.Writer
    DryRun bool
}

func NewSession(in io.Reader, out io.Writer, errs io.Writer) *Session{
    return &Session{
        Stdin: in,
        Stdout: out,
        Stderr: errs,
        DryRun: false,
    }
}

func (s *Session) Run() {
    fmt.Fprint(s.Stdout, "> ")
    input := bufio.NewScanner(s.Stdin)
    for input.Scan(){
        line := input.Text()
        cmd, err := CmdFromString(line)
        if err != nil {
            fmt.Fprint(s.Stdout, "> ")
            continue
        }
        if s.DryRun {
            fmt.Fprintf(s.Stdout, "%s\n> ", line)
            continue
        }
        output, err := cmd.CombinedOutput()
        if err != nil {
            fmt.Fprintln(s.Stderr, "error:", err)
        }
        fmt.Fprintf(s.Stdout, "%s>", output)
    }
    fmt.Fprint(s.Stdout, "\nUntil next time, earthling!")
}