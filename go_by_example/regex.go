package main

import (
	"fmt"
	"regexp"
)

func main() {

	// string pattern directly
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Generate a reg-exp struct
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))
}
