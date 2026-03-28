package main

import "strings"

func Encode(message string) string {
	vowels := "aeiouAEIOU"

	// Create a map for vowel replacements
	vowelReplace := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		vowelReplace[v] = v + 2
	}

	// Map function that swaps case and replaces vowels
	return strings.Map(func(r rune) rune {
		// Swap case: if lowercase, make uppercase; if uppercase, make lowercase
		if r >= 'a' && r <= 'z' {
			r -= 'a' - 'A'
		} else if r >= 'A' && r <= 'Z' {
			r += 'a' - 'A'
		}

		// Replace vowels with letter 2 positions ahead
		if replacement, ok := vowelReplace[r]; ok {
			return replacement
		}
		return r
	}, message)
}
