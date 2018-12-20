package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)

	ints := []int{4, 3, 2, 1}
	sort.Ints(ints)

	fmt.Println(strs, ints)
	fmt.Println("Are strings sorted? ", sort.StringsAreSorted(strs))
}
