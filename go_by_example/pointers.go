package main

import "fmt"

// Passed by value
func zeroval(input int) {
	input = 0
}

// passed by reference
func zeroptr(input *int) {
	*input = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Dereferencing a pointer (like in C)
	fmt.Println("ptr:", &i)
}
