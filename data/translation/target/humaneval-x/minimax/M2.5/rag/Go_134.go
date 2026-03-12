package main

import (
	"strings"
	"unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
	// Split the string by spaces and get the last element
	parts := strings.Split(txt, " ")
	last := parts[len(parts)-1]
	
	// Check if length is exactly 1 and it's a letter
	if len(last) == 1 {
		r := rune(last[0])
		return unicode.IsLetter(r)
	}
	return false
}

// For testing (can be run standalone)
func main() {
	// Test cases
	testCases := []string{
		"apple pie",
		"apple pi e",
		"apple pi e ",
		"",
	}
	
	for _, tc := range testCases {
		println(tc, "->", CheckIfLastCharIsALetter(tc))
	}
}
