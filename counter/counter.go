package counter

import (
	"bufio"
	"io"
)

type Counter struct {
	Lines int
	Input io.Reader
}

func (c *Counter) Count() {
	input := bufio.NewScanner(c.Input)

	for input.Scan() {
		c.Lines++
	}
}

func (c *Counter) GetLines() int {
	return c.Lines
}

func NewCounter(r io.Reader) *Counter {
	return &Counter{0, r}
}

func CountLines(r io.Reader) int {
	c := NewCounter(r)
	c.Count()
	return c.GetLines()
}
