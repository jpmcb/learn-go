package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {

		// If missing beginning of url, append to string
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		// Also prints the status
		fmt.Println(resp.Status)
	}
}
