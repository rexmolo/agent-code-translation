package main

import (
	"strings"
	"unicode"
)

// Encode swaps the case of all letters and replaces vowels with the letter 
// that appears 2 places ahead of it in the alphabet.
func Encode(message string) string {
	var result strings.Builder
	result.Grow(len(message))

	const vowels = "aeiouAEIOU"

	for _, r := range message {
		// First, determine the case-swapped character.
		swappedChar := r
		if unicode.IsLower(r) {
			swappedChar = unicode.ToUpper(r)
		} else if unicode.IsUpper(r) {
			swappedChar = unicode.ToLower(r)
		}

		// Next, check if the case-swapped character is a vowel.
		if strings.ContainsRune(vowels, swappedChar) {
			// If it is, replace it with the character 2 places ahead.
			result.WriteRune(swappedChar + 2)
		} else {
			// Otherwise, use the case-swapped character as is.
			result.WriteRune(swappedChar)
		}
	}

	return result.String()
}
