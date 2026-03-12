package main

import (
	"strings"
	"unicode"
)

// CheckIfLastCharIsALetter returns True if the last character
// of a given string is an alphabetical character and is not
// a part of a word, and False otherwise.
// Note: "word" is a group of characters separated by space.
func CheckIfLastCharIsALetter(txt string) bool {
	// Split string by space
	parts := strings.Split(txt, " ")
	// Get the last element
	lastPart := parts[len(parts)-1]
	// Check if length is 1 and it's a letter
	if len(lastPart) == 1 {
		r := []rune(lastPart)[0]
		return unicode.IsLetter(r)
	}
	return false
}