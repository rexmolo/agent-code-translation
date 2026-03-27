package main

import (
	"fmt"
	"strings"
)

func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")
	lastPart := parts[len(parts)-1]

	if len(lastPart) == 1 {
		ch := lastPart[0]
		// Check if it's an alphabetical character (a-z or A-Z)
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			return true
		}
	}
	return false
}

func main() {
	// Test cases
	fmt.Println(CheckIfLastCharIsALetter("apple pie"))   // false
	fmt.Println(CheckIfLastCharIsALetter("apple pi e"))  // true
	fmt.Println(CheckIfLastCharIsALetter("apple pi e ")) // false
	fmt.Println(CheckIfLastCharIsALetter(""))            // false
}