package main

import (
	hello "github.com/igenexxx/hello/greetings"
	"os"
)

func main() {
	hello.PrintTo(os.Stdout)
}
