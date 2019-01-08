// Implements reversing a Vigenère cipher.
// Given a key and a cipher text, returns the decrypted plain text

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		panic("Usage: decode {key} {message}")
	}

	key := strings.ToUpper(os.Args[1])
	code := strings.ToUpper(os.Args[2])

	var res string

	// Go through each char of the key
	for i, c := range key {
		temp := 65

		// Increment the temp placeholder starting at 'A'
		for {
			// .. until we find the corisponding k ey
			if c == rune(code[i]) {
				break
			} else {
				// .. otherwise, increment the temp placeholder and current char
				temp++
				c++
				if c > 90 {
					// For wrapping of the ceaser cypher
					c = (c % 90) + 64
				}
			}
		}

		res += string(temp)
	}

	fmt.Println(res)
}
