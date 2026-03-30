package main

import (
	"strings"
	"unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
	// Split the string by space and get the last element
	parts := strings.Split(txt, " ")
	last := parts[len(parts)-1]

	// Check if length is exactly 1 and it's an alphabetical character
	if len(last) == 1 {
		ch := rune(last[0])
		lower := unicode.ToLower(ch)
		if lower >= 'a' && lower <= 'z' {
			return true
		}
	}

	return false
}