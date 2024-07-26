package counter_test

import (
	"bytes"
	"github.com/igenexxx/counter"
	"testing"
)

func TestCounter_CountShouldBeValid(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	buf.WriteString("Hello\nWorld\n")

	lines := counter.CountLines(buf)

	if lines != 2 {
		t.Errorf("Expected 2, got %d", lines)
	}
}
