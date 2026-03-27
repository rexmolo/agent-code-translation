package main

import (
	"strings"
	"unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
	// Split by space and get the last element
	parts := strings.Split(txt, " ")
	lastPart := parts[len(parts)-1]

	// Check if length is exactly 1 and is a letter
	if len(lastPart) == 1 {
		ch := rune(lastPart[0])
		if unicode.IsLetter(ch) {
			return true
		}
	}
	return false
}