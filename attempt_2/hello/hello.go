package hello

import (
	"fmt"
	"io"
)

func PrintTo(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s\n", name)
}
