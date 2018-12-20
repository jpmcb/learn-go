package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// Creates a struct
	fmt.Println(person{"Bob", 20})

	// Specific fields
	fmt.Println(person{name: "kevin", age: 31})

	// Fields not included are zero valued
	fmt.Println(person{name: "Joan"})

	// Get the pointer for the struct itself
	fmt.Println(&person{name: "Sean", age: 55})

	// Get fields with a dot operator
	s := person{name: "Ann", age: 22}
	fmt.Println(s.name)

	// Dots work with dereferencing pointers
	sp := &s
	fmt.Println(sp.age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(*sp)
}
