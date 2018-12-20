package main

import "fmt"

// func, name, arguments, return type, code block
func plus(a int, b int) int {
	return a + b
}

// Can omit being specific about each argument type
func plusPlus(a, b, c int) int {
	return a + b + c
}

// Mulitple return values
func vals() (int, int) {
	return 3, 7
}

func main() {
	res := plus(1, 2)
	fmt.Println("1 + 2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", res)

	// Multipe return vals
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
