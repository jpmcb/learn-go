package main

import "fmt"

func main() {

	// Trivial for looping structure
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// C style looping structure
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// Breaking a looping structure
	for {
		fmt.Println("inside loop")
		break
	}

	// C style loop with continues
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
