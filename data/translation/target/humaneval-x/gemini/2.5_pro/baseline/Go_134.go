package main

import (
	"fmt"
	"strings"
)

// CheckIfLastCharIsALetter returns True if the last character
// of a given string is an alphabetical character and is not
// a part of a word, and False otherwise.
// Note: "word" is a group of characters separated by space.
func CheckIfLastCharIsALetter(txt string) bool {
	// The Python code's logic is to check the substring after the last space.
	// `strings.LastIndex` is more efficient than `strings.Split` for this purpose.
	lastSpaceIndex := strings.LastIndex(txt, " ")
	check := txt[lastSpaceIndex+1:]

	// The condition is that this substring must be a single alphabetical character.
	if len(check) != 1 {
		return false
	}

	char := check[0] // It's safe to access the first byte as we've checked the length.

	// This check for an ASCII letter directly mimics the behavior of Python's
	// `97 <= ord(check.lower()) <= 122`.
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
}

// main function to demonstrate the usage with examples from the problem description.
func main() {
	examples := []struct {
		input    string
		expected bool
	}{
		{"apple pie", false},
		{"apple pi e", true},
		{"apple pi e ", false},
		{"", false},
	}

	for _, ex := range examples {
		result := CheckIfLastCharIsALetter(ex.input)
		fmt.Printf("CheckIfLastCharIsALetter(\"%s\") -> %v (Expected: %v)\n", ex.input, result, ex.expected)
	}
}
