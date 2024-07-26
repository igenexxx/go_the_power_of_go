package hi

import (
	"fmt"
	"io"
	"os"
)

type Printer struct {
	Output io.Writer
}

func (p *Printer) Print() {
	fmt.Fprintln(p.Output, "Hi, world!")
}

func NewPrinter() *Printer {
	return &Printer{Output: os.Stdout}
}

func Main() {
	p := NewPrinter()
	p.Print()
}
