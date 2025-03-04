package shell

import (
	"errors"
	"io"
	"os/exec"
	"strings"
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
}

func NewSession(in io.Reader, out io.Writer, errs io.Writer) *Session{
    return &Session{
        Stdin: in,
        Stdout: out,
        Stderr: errs,
    }
}