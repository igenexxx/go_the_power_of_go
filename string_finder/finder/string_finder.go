package finder

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type finder struct {
	input  io.Reader
	output io.Writer
}

type option func(*finder) error

func (f *finder) FindString(substring string) {
	scanner := bufio.NewScanner(f.input)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, substring) {
			_, err := fmt.Fprint(f.output, line)
			if err != nil {
				log.Printf("error writing to output: %v", err)
			}
		}
	}
}

func New(opts ...option) (*finder, error) {
	f := &finder{
		input:  os.Stdin,
		output: os.Stdout,
	}

	for _, opt := range opts {
		if err := opt(f); err != nil {
			return nil, err
		}
	}

	return f, nil
}

func WithInput(input io.Reader) option {
	return func(f *finder) error {
		if input == nil {
			return fmt.Errorf("input cannot be nil")
		}

		f.input = input
		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(f *finder) error {
		if output == nil {
			return fmt.Errorf("output cannot be nil")
		}

		f.output = output
		return nil
	}
}

func Main() {
	substr := os.Args[1]
	if len(substr) == 0 {
		log.Fatalln("Mandatory argument is absent")
	}

	f, err := New()
	if err != nil {
		panic(err)
	}

	f.FindString(substr)
}
