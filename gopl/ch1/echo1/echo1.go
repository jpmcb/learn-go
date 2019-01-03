package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	// Loop init. Build the string ground up
	for i := 1; i < len(os.Args); i++ {
		s += (sep + os.Args[i])
		sep = " "
	}

	fmt.Println(s)
}
