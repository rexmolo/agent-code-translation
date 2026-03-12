package main

import "strings"

// RemoveVowels takes a string and returns a new string with all vowels removed.
func RemoveVowels(text string) string {
	var result strings.Builder
	for _, r := range text {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			// It's a vowel, so we skip it.
		default:
			// It's not a vowel, so we append it to our result.
			result.WriteRune(r)
		}
	}
	return result.String()
}
