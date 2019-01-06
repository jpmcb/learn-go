package main

import (
	"fmt"
	"os"
	"strings"
)

// Prints the entire list of given arguments including inducing program
func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
