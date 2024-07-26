package hi_test

import (
	"bytes"
	"github.com/igenexxx/hello/hi"
	"testing"
)

func TestHi_CheckPrintMessage(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	want := "Hi, world!\n"

	h := hi.NewPrinter()
	h.Output = buf
	h.Print()

	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
