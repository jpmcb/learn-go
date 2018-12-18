package main

import "fmt"

func main() {

	// An infered string
	var a = "initial"
	fmt.Println(a)

	// A few ints
	var b, c int = 1, 2
	fmt.Println(b, c)

	// A bool variable
	var d bool = true
	fmt.Println(d)

	// Initializes to empty zero value
	var e int 
	fmt.Println(e)

	// The shorthand for creating a string infered
	f := "short"
	fmt.Println(f)
}