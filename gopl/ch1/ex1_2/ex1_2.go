package main

import (
	"fmt"
	"os"
	"strconv"
)

// Includes the zeroed index of each argument
func main() {
	var s string
	var sep string

	for i, arg := range os.Args {
		fmt.Print(string(i))
		s += sep + strconv.Itoa(i) + " " + arg
		sep = " "
	}

	fmt.Println(s)
}
