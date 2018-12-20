package main

import "fmt"

func main() {

	// Sum of numbers in slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// Range returns both the index and the element
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range used on map (key value pairs)
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Can just iterate keys
	for k := range m {
		fmt.Printf("Key: %s\n", k)
	}

	// Iterate over index / unicode
	for i, c := range "go" {
		fmt.Printf("Index: %d Uni: %d\n", i, c)
	}

}
