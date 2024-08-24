package hello_test

import (
	"bytes"
	"github.com/igenexxx/hello"
	"testing"
)

func TestPrintTo_PrintsHelloMessageToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)

	p := hello.New()
	p.Output = buf
	p.Print()

	want := "Hello, world!\n"
	got := buf.String()

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
