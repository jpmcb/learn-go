// Implements the rot13 cipher.
// Rotates each char by 13 places. Can also be used to reverse a rotation.

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Usage: decode {cipher-text}")
	}

	code := os.Args[1]

	var res string

	// Go through each char of the input
	for _, c := range code {
		temp := c

		// If letter is lower case, rotate and wrap
		if c >= 97 && c <= 122 {
			temp += 13
			if temp > 122 {
				temp = (temp % 122) + 96
			}

			// If letter is upper case, rotate and wrap
		} else if c >= 65 && c <= 90 {
			temp += 13
			if temp > 90 {
				temp = (temp % 90) + 64
			}
		}

		res += string(temp)
	}

	fmt.Println(res)
}
