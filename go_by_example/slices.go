package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// Similar to arrays
	s[0] = "one"
	s[1] = "two"
	s[2] = "three"

	fmt.Println("set element:", s)
	fmt.Println("get ele:", s[2])
	fmt.Println("length:", len(s))

	// Can use the append function
	s = append(s, "four")
	s = append(s, "five", "six")
	fmt.Println("appended:", s)

	// Copy elements into new slice
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy:", c)

	// Slice operator 2 - 5 not including 5
	l := s[2:5]
	fmt.Println("sl1:", l)

	// Up to but not including 5th index
	l = s[:5]
	fmt.Println("sl2:", l)

	// Up from and including 2nd index
	l = s[2:]
	fmt.Println("sl3:", l)

	// slice in one liner
	t := []string{"a", "b", "c"}
	fmt.Println("one liner:", t[1:2])

	// 2 Demensional slices (with variable lengths)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		inner := i + 1
		twoD[i] = make([]int, inner)
		for j := 0; j < inner; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
