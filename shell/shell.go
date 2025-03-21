package shell

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
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
	Stdin      io.Reader
	Stdout     io.Writer
	Stderr     io.Writer
	DryRun     bool
	Transcript io.Writer
}

func NewSession(in io.Reader, out io.Writer, errs io.Writer) *Session {
	return &Session{
		Stdin:  in,
		Stdout: out,
		Stderr: errs,
        Transcript: io.Discard,
		DryRun: false,
	}
}

func (s *Session) Run() {
	stdout := io.MultiWriter(s.Stdout, s.Transcript)
	stderr := io.MultiWriter(s.Stderr, s.Transcript)
	fmt.Fprintf(stdout, "> ")
	input := bufio.NewScanner(s.Stdin)
	for input.Scan() {
		line := input.Text()
		cmd, err := CmdFromString(line)
		if err != nil {
			fmt.Fprintf(stdout, "> ")
			continue
		}
		if s.DryRun {
			fmt.Fprintf(stdout, "%s\n> ", line)
			continue
		}
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(stderr, "error:", err)
		}
		fmt.Fprintf(stdout, "%s>", output)
	}
	fmt.Fprintln(stdout, "\nUntil next time, earthling!")
}

func Main() {
	session := NewSession(os.Stdin, os.Stdout, os.Stderr)
	transcript, err := os.Create("transcript.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer transcript.Close()
	session.Transcript = transcript
	session.Run()
	fmt.Println("[output file is 'transcript.txt']")
}
