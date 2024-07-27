package main

import (
	"github.com/igenexxx/string-finder-go/finder"
	"os"
)

func main() {
	wordToSearch := os.Args[1]
	if wordToSearch == "" {
		panic("No required argument")
	}

	f, err := finder.New()
	if err != nil {
		panic("error when create finder")
	}

	f.FindString(wordToSearch)
}
