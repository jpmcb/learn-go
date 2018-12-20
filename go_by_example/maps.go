package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["n1"] = 4
	m["n2"] = 7

	fmt.Println("map:", m)
	fmt.Println("Get ele:", m["n1"])
	fmt.Println("map len:", len(m))

	// Remove an element from a map
	delete(m, "n2")
	fmt.Println("map with 1 ele now:", m)

	// 2 return values. The underscore ( _ ) is a blank identifier
	// since we only need the second to check if present in map
	_, present := m["n2"]
	fmt.Println("Present?", present)

	// One liner map
	one := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("new map:", one)
}
