package main

import (
	"strings"
	"unicode"
)

// CheckIfLastCharIsALetter returns true if the last character
// of a given string is an alphabetical character and is not
// a part of a word (i.e., the last "word" is a single character).
func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")

	// Handle empty string case
	if len(parts) == 0 {
		return false
	}

	last := parts[len(parts)-1]

	// Must be exactly one character
	if len(last) != 1 {
		return false
	}

	// Check if it's a letter (a-z or A-Z)
	return unicode.IsLetter(rune(last[0]))
}

func main() {
	// Test cases
	println(CheckIfLastCharIsALetter("apple pie"))      // false
	println(CheckIfLastCharIsALetter("apple pi e"))     // true
	println(CheckIfLastCharIsALetter("apple pi e "))   // false
	println(CheckIfLastCharIsALetter(""))              // false
}