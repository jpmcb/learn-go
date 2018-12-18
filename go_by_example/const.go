package main 

import "fmt"
import "math"

const s string = "constant"

func main() {

	// In global scoping
	fmt.Println(s)

	const n = 50000000

	// the const expression performs arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// Cast the ambiguous const
	fmt.Println(int64(n))

	// Sin expects a int64, go will infer this
	fmt.Println(math.Sin(n))
}