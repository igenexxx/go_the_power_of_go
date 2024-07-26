package counter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input  io.Reader
	output io.Writer
}

func New() *counter {
	return &counter{
		input: os.Stdin,
	}
}

func (c *counter) Lines() int {
	var lines int
	input := bufio.NewScanner(c.input)

	for input.Scan() {
		lines++
	}

	return lines
}

func Count() int {
	c := New()
	return c.Lines()
}

func PrintLinesCount() {
	lines := Count()
	fmt.Println(lines)
}

type option func(*counter) error

func NewCounter(opts ...option) (*counter, error) {
	c := &counter{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithInput(input io.Reader) option {
	return func(c *counter) error {
		if input == nil {
			return errors.New("nil input reader")
		}

		c.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(c *counter) error {
		if output == nil {
			return errors.New("nil output reader")
		}

		c.output = output
		return nil
	}
}

func Main() {
	c, err := NewCounter()
	if err != nil {
		panic(err)
	}

	fmt.Println(c.Lines())
}
