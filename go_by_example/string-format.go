package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	// Prints the field names
	fmt.Printf("%+v\n", p)

	// Prints the source code that would print this
	fmt.Printf("%#v\n", p)

	// Prints the type of the value
	fmt.Printf("%T\n", p)

	// Prints a bool
	fmt.Printf("%t\n", true)

	// Prints a digit
	fmt.Printf("%d\n", 123)

	// Prints bin number
	fmt.Printf("%b\n", 14)

	// prints hex encoding
	fmt.Printf("%x\n", 456)

	// Prints a char (given by specific char or int value)
	fmt.Printf("%c\n", 33)

	// Formatting for floats
	fmt.Printf("%f\n", 78.9)

	// sci notation
	fmt.Printf("%e\n", 12340000.0)
	fmt.Printf("%E\n", 12340000.0)

	// string printing
	fmt.Printf("%s\n", "\"string\"")

	// Double quoting
	fmt.Printf("%q\n", "\"string\"")

	// Encodes string in hex val
	fmt.Printf("%x\n", "hex this")

	// Prints a memory address from pointer
	fmt.Printf("%p\n", &p)

	// Specify width of the output
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
}
