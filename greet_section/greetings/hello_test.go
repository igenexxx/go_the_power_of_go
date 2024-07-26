package hello_test

import (
	"bytes"
	hello "github.com/igenexxx/hello/greetings"
	"testing"
)

func TestPrintTo_CheckPrintMessage(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	want := "Hello, world!\n"

	hello.PrintTo(buf)

	got := buf.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
