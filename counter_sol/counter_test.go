package counter_test

import (
	"bytes"
	"github.com/igenexxx/counter"
	"github.com/rogpeppe/go-internal/testscript"
	"log"
	"os"
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

func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := counter.NewCounter(counter.WithInputFromArgs(args))
	if err != nil {
		t.Fatal(err)
	}

	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3")
	c, err := counter.NewCounter(counter.WithInput(inputBuf), counter.WithInputFromArgs([]string{}))
	if err != nil {
		t.Fatal(err)
	}

	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}

func TestMain(m *testing.M) {
	os.Exit(
		testscript.RunMain(
			m,
			map[string]func() int{
				"counter": counter.Main,
			},
		),
	)
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/scripts",
	})
}
