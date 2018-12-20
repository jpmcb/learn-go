package main

import "fmt"

func main() {
	var A [5]int
	fmt.Println("empty array:", A)

	A[4] = 100
	fmt.Println("set an element:", A)
	fmt.Println("get an element:", A[4])

	// Length of array
	fmt.Println("Len:", len(A))

	// One liner
	B := [5]int{1, 2, 3, 4, 5}
	fmt.Println(B)

	// Messing with a 2d array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
