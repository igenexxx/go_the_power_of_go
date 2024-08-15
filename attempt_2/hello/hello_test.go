package hello_test

import (
	"github.com/igenexxx/hello"
	"testing"
)

func TestPrint_PrintsMessageToTerminal(t *testing.T) {
	t.Parallel()
	hello.Print()
}
