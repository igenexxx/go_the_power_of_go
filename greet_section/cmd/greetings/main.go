package main

import (
	"fmt"
	hello "github.com/igenexxx/hello/greetings"
	"github.com/igenexxx/hello/hi"
	"net/http"
	"os"
	"time"
)

func main() {
	hello.PrintTo(os.Stdout)

	hi.Main()

	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Println(resp.Status)

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	res, err := client.Get("https://google.com")
	if err != nil {
		fmt.Errorf("error: %v", err)
	}

	fmt.Println(res.Status)
}
