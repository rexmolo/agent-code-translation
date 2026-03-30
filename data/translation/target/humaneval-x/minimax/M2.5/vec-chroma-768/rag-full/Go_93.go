package main

import (
	"strings"
	"unicode"
)

func Encode(message string) string {
	vowels := "aeiouAEIOU"

	// Build a map for vowel replacements (each vowel -> vowel + 2 positions)
	vowelsReplace := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		vowelsReplace[v] = rune(int(v) + 2)
	}

	// Use strings.Map to transform each character
	return strings.Map(func(r rune) rune {
		// Swap case: lowercase -> uppercase, uppercase -> lowercase
		if r >= 'a' && r <= 'z' {
			r = unicode.ToUpper(r)
		} else if r >= 'A' && r <= 'Z' {
			r = unicode.ToLower(r)
		}

		// Replace vowel with the character 2 positions ahead
		if repl, ok := vowelsReplace[r]; ok {
			return repl
		}
		return r
	}, message)
}
