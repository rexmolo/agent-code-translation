package main

import (
	"unicode"
)

// Solve processes a string according to specific rules.
// If a character is a letter, its case is reversed.
// If the string contains no letters, the entire string is reversed.
// Otherwise, non-letter characters are kept as they are.
func Solve(s string) string {
	runes := []rune(s)
	hasLetter := false

	for i, r := range runes {
		if unicode.IsLetter(r) {
			hasLetter = true
			if unicode.IsUpper(r) {
				runes[i] = unicode.ToLower(r)
			} else {
				runes[i] = unicode.ToUpper(r)
			}
		}
	}

	if !hasLetter {
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}

	return string(runes)
}
