package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}

		// Uses copy instaed of buffer to place body contents into stdio
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Printf("err: %v\n", err)
			os.Exit(1)
		}
	}
}
