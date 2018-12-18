package main

import "fmt"

func main() {
	var input = 6
	var res string

	// Fizzbuzz implementation using branching if / else blocks
	if input%2 == 0 && input%3 == 0 {
		res = "FizzBuzz"
	} else if input%2 == 0 {
		res = "Fizz"
	} else if input%3 == 0 {
		res = "Buzz"
	} else {
		res = "Foo"
	}

	fmt.Println(res)
}
