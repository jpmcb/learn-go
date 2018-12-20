package main

import "fmt"

// A Simple function
func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, " : ", i)
	}
}

func main() {

	// Blocking sequential function
	f("direct")

	// Concurrent goroutine
	go f("goroutine")

	// Should concurrently run with the above f call
	go func(msg string) {
		fmt.Println(msg)
	}("going anon")

	fmt.Scanln()
	fmt.Println("done")
}
