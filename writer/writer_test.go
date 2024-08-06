package writer_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/igenexxx/writer"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteToFile_WritesGivenDataToFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "write_test.txt")

	want := []byte{1, 2, 3}
	err := writer.WriteToFile(path, want)
	if err != nil {
		t.Fatal(err)
	}

	stat, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}

	perm := stat.Mode().Perm()
	if perm != os.FileMode(0600) {
		t.Errorf("want file mode 0600, got O%o", perm)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestWriteToFile_ReturnsErrorWhenFileDoesNotExist(t *testing.T) {
	t.Parallel()
	path := filepath.Join("not_exist", "write_test.txt")
	err := writer.WriteToFile(path, nil)
	if err == nil {
		t.Fatal("want error when directory does not exist")
	}
}

func TestWriteToFile_RewriteExistingFile(t *testing.T) {
	t.Parallel()

	path := filepath.Join(t.TempDir(), "write_test.txt")

	err := os.WriteFile(path, []byte("Hello World"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	want := []byte{1, 2, 3}

	err = writer.WriteToFile(path, want)
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
