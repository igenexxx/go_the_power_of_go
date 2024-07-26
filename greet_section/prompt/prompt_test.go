package prompt_test

import (
	"bytes"
	"github.com/igenexxx/hello/prompt"
	"testing"
)

func TestUserInput_HasText(t *testing.T) {
	t.Parallel()
	writeBuf := new(bytes.Buffer)
	readBuf := new(bytes.Buffer)

	prompt.UserInput(readBuf, writeBuf, "What is your name?")
	writeBuf.WriteString("John\n")

	questionWant := "What is your name?"
	answerWant := "Hello, John!"
	questionGot := writeBuf.String()
	answerGot := readBuf.String()

	if questionGot != questionWant {
		t.Errorf("got %q, want %q", questionGot, questionWant)
	}

	if answerGot != answerWant {
		t.Errorf("got %q, want %q", answerGot, answerWant)
	}
}
