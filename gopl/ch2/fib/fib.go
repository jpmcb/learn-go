package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: fib <int>")
	}

	i, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println(fib(i))
}

func fib(seq int) int {
	x, y := 0, 1
	for i := 0; i < seq; i++ {

		// Uses tuple unpacking to increment to fib sequence
		x, y = y, x+y
	}
	return x
}
