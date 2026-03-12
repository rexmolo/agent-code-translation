package main

import (
	"strings"
)

func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")
	if len(parts) == 0 {
		return false
	}

	check := parts[len(parts)-1]

	if len(check) == 1 {
		ch := rune(check[0])
		return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
	}
	return false
}

func main() {
	// Test cases
	println(CheckIfLastCharIsALetter("apple pie"))      // false
	println(CheckIfLastCharIsALetter("apple pi e"))    // true
	println(CheckIfLastCharIsALetter("apple pi e "))   // false
	println(CheckIfLastCharIsALetter(""))              // false
}
