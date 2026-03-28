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

	last := parts[len(parts)-1]
	if len(last) != 1 {
		return false
	}

	// Get the lowercase version of the character and check if it's a letter
	r := rune(last[0])
	lowerR := unicode.ToLower(r)
	if lowerR >= 'a' && lowerR <= 'z' {
		return true
	}

	return false
}