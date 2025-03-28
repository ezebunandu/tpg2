package match

import (
	"io"
	"os"
    "fmt"
    "bufio"
    "strings"
    "errors"
)

type matcher struct {
    input io.Reader
    output io.Writer
    text string
}

type option func(*matcher) error

func WithInput(input io.Reader) option {
    return func(m *matcher) error {
        if input == nil {
            return errors.New("nil input reader")
        }
        m.input = input
        return nil
    }
}

func WithOutput(output io.Writer) option {
    return func(m *matcher) error {
        if output == nil {
            return errors.New("nil output writer")
        }
        m.output = output
        return nil
    }
}

func WithSearchStringFromArgs(args []string) option {
    return func(m *matcher) error {
        if len(args) < 1 {
            return nil
        }
        m.text = args[0]
        return nil
    }
}


func NewMatcher(opts ...option) (*matcher, error) {
    m := &matcher{
        input: os.Stdin,
        output: os.Stdout,
    }
    for _, opt := range opts {
        err := opt(m)
        if err != nil {
            return nil, err
        }
    }
    return m, nil
}

func (m *matcher) PrintMatchingLines() {
    input := bufio.NewScanner(m.input)
    for input.Scan() {
        if strings.Contains(input.Text(), m.text) {
            fmt.Fprintln(m.output, input.Text())
        }
    }
}

func Main() int {
    m, err := NewMatcher(
        WithSearchStringFromArgs(os.Args[1:]),
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        return 1
    }
    m.PrintMatchingLines()
    return 0
}