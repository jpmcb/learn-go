package main

import "fmt"

// gets the n-th number in the fib sequence
func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	// Prints the first 7 numbers of the fib sequence
	fmt.Println(fib(0))
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
}
