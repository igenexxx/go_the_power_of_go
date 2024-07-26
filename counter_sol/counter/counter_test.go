package counter_test

import (
	"bytes"
	"github.com/igenexxx/counter/counter"
	"log"
	"testing"
)

func TestCounter_CountShouldBeValid(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("one\ntwo\nthree")

	c, err := counter.NewCounter(counter.WithInput(inputBuf))
	if err != nil {
		log.Fatalln(err)
	}

	want := 3
	got := c.Lines()

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
