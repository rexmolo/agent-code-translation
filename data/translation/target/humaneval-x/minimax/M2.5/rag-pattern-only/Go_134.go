package main

import (
    "strings"
    "unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")
	if len(parts) == 0 {
		return false
	}

	lastPart := parts[len(parts)-1]

	if len(lastPart) != 1 {
		return false
	}

	// Check if it's a letter (a-z or A-Z) using unicode.IsLetter
	runeVal := []rune(lastPart)[0]
	return unicode.IsLetter(runeVal)
}

func main() {
	// Test cases
	println(CheckIfLastCharIsALetter("apple pie"))   // false
	println(CheckIfLastCharIsALetter("apple pi e")) // true
	println(CheckIfLastCharIsALetter("apple pi e ")) // false
	println(CheckIfLastCharIsALetter(""))            // false
}