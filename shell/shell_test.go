package shell_test

import (
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