package prompt

import (
	"bufio"
	"fmt"
	"io"
)

func UserInput(r io.Reader, w io.Writer, q string) (string, error) {
	reader := bufio.NewReader(r)
	fmt.Fprint(w, q)

	prompt, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return prompt, nil
}
