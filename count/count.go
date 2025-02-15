package count

import (
	"bufio"
	"io"
	"os"
)

type Counter struct {
	Input io.Reader
}

func NewCounter() *Counter {
	return &Counter{
		Input: os.Stdin,
	}
}

func (c *Counter) Lines() int {
    lines := 0
    input := bufio.NewScanner(c.Input)
    for input.Scan(){
        lines++
    }
    return lines
}

func Main(){
    NewCounter().Lines()
}