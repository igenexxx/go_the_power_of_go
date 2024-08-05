package counter_test

import (
	"bytes"
	"github.com/google/go-cmp/cmp"
	"github.com/igenexxx/counter"
	"github.com/rogpeppe/go-internal/testscript"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/scripts",
	})
}

func TestMain(m *testing.M) {
	os.Exit(
		testscript.RunMain(
			m,
			map[string]func() int{
				"count": counter.Main,
			},
		),
	)
}

//func TestMain(m *testing.M) {
//	os.Exit(
//		testscript.RunMain(
//			m,
//			map[string]func() int{
//				"lines": counter.MainLines,
//				"words": counter.MainWords,
//			},
//		),
//	)
//}

func TestLinesCounter_CountShouldBeValid(t *testing.T) {
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

func TestLinesCounterWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
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

func TestLinesCounterWithInputFromArgs_IgnoresEmptyArgs(t *testing.T) {
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

func TestWordsCounter_CountShouldBeValid(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("one line before the second on\nsecond line\nthird line here")

	c, err := counter.NewCounter(counter.WithInput(inputBuf))
	if err != nil {
		log.Fatalln(err)
	}

	want := 11
	got := c.Words()

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWriteToFile_WritesGivenDataToFile(t *testing.T) {
	t.Parallel()
	path := "testdata/write_test.txt"
	want := []byte{1, 2, 3}
	err := counter.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}
