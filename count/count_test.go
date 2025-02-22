package count_test

import (
	"bytes"
	"testing"
    "os"

	"github.com/ezebunandu/count"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestLines_CountsLinesInInput(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(count.WithInput(inputBuf))
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWords_CountsWordsInInput(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("one two three four\n")
	c, err := count.NewCounter(count.WithInput(inputBuf))
	if err != nil {
		t.Fatal(err)
	}
	want := 4
	got := c.Words()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}

	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := count.NewCounter(
		count.WithInput(inputBuf),
		count.WithInputFromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"count": count.Main,
	}))
}
