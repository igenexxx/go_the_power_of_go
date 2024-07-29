package counter

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type counter struct {
	input  io.Reader
	output io.Writer
	files  []*os.File
}

func (c *counter) Close() error {
	for _, file := range c.files {
		if err := file.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (c *counter) Lines() int {
	var lines int
	input := bufio.NewScanner(c.input)

	for input.Scan() {
		lines++
	}

	return lines
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

func WithInputFromArgs(args []string) option {
	return func(c *counter) error {
		if len(args) == 0 {
			return nil
		}

		readers := make([]io.Reader, len(args))
		for i, filename := range args {
			f, err := os.Open(filename)
			if err != nil {
				return err
			}
			c.files = append(c.files, f)
			readers[i] = f
		}

		c.input = io.MultiReader(readers...)

		return nil
	}
}

func MainLines() int {
	c, err := NewCounter(WithInputFromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer c.Close()

	fmt.Println(c.Lines())
	return 0
}

func MainWords() int {
	c, err := NewCounter(WithInputFromArgs(os.Args[1:]))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer c.Close()

	fmt.Println(c.Words())
	return 0
}

func (c *counter) Words() int {
	var words int
	input := bufio.NewScanner(c.input)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		words++
	}

	return words
}

func Main() int {
	lineMode := flag.Bool("lines", false, "Count lines, not words")
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-lines] [files...]\n\n", os.Args[0])
		fmt.Print("  Counts words (or lines) from stdin (or files)\n\n")
		fmt.Println("Flags: ")
		flag.PrintDefaults()
	}
	flag.Parse()

	c, err := NewCounter(WithInputFromArgs(flag.Args()))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	if *lineMode {
		fmt.Println(c.Lines())
	} else {
		fmt.Println(c.Words())
	}

	return 0
}
