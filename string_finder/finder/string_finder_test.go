package finder_test

import (
	"bytes"
	"github.com/igenexxx/string-finder-go/finder"
	"log"
	"testing"
)

func TestFindString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		content  string
		input    string
		expected string
	}{
		{"empty string", "", "", ""},
		{"single word", "this is hello string", "hello", "this is hello string"},
		{"multiple words", "that string contain hello world so", "hello world", "that string contain hello world so"},
		{"separators", "this string contain hello, world, and more", "hello world", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputBuf := bytes.NewBufferString(tt.content)
			outputBuf := new(bytes.Buffer)
			f, err := finder.New(finder.WithInput(inputBuf), finder.WithOutput(outputBuf))
			if err != nil {
				log.Fatalln(err)
			}

			f.FindString(tt.input)

			got := outputBuf.String()

			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
