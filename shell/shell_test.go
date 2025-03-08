package shell_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/ezebunandu/shell"
	"github.com/google/go-cmp/cmp"
)

func TestCmdFromString_CreatesExpectedCmd(t *testing.T){
    t.Parallel()
    input := "/bin/ls -l main.go"
    want := []string{"/bin/ls", "-l", "main.go"}
    cmd, err := shell.CmdFromString(input)
    if err != nil{
        t.Fatal(err)
    }
    got := cmd.Args
    if !cmp.Equal(want, got) {
        t.Error(cmp.Diff(want, got))
    }
}

func TestCmdFromString_ErrorsOnEmptyInput(t *testing.T){
    t.Parallel()
    _, err := shell.CmdFromString("")
    if err == nil {
        t.Fatal("want error on empty input, got nil!")
    }
}

func TestNewSession_CreatesExpectedSession(t *testing.T){
    t.Parallel()
    want := shell.Session{
        Stdin: os.Stdin,
        Stdout: os.Stdout,
        Stderr: os.Stderr,
        DryRun: false,
        Transcript: io.Discard,
    }
    got := *shell.NewSession(os.Stdin, os.Stdout, os.Stderr)

    if want != got {
        t.Errorf("want %#v, got %#v", want, got)
    }
}

func TestRun_ProducesExpectedOutput(t *testing.T){
    t.Parallel()
    stdin := strings.NewReader("echo hello\n\n")
    stdout := new(bytes.Buffer)
    session := shell.NewSession(stdin, stdout, io.Discard)
    session.DryRun = true
    session.Run()
    want := "> echo hello\n> > \nUntil next time, earthling!\n"
    got := stdout.String()
    if !cmp.Equal(want, got){
        t.Error(cmp.Diff(want, got))
    }
}

func TestRunProducesExpectedTranscript(t *testing.T) {
	t.Parallel()
    in := strings.NewReader("echo hello\n\n")
    transcript := new(bytes.Buffer)
    session := shell.NewSession(in, io.Discard, io.Discard)
    session.DryRun = true
    session.Transcript = transcript
    session.Run()
    want := "> echo hello\n> > \nUntil next time, earthling!"
    got := transcript.String()
    if !cmp.Equal(want, got) {
        cmp.Diff(want, got)
    }
}