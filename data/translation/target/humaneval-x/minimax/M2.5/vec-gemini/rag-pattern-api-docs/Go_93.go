package main

import (
	"strings"
	"unicode"
)

func Encode(message string) string {
	// Create vowel replacement map
	vowels := "aeiouAEIOU"
	vowelsReplace := make(map[rune]rune)
	for _, v := range vowels {
		vowelsReplace[v] = rune(int(v) + 2)
	}

	// Swap case using strings.Map
	swapped := strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		return unicode.ToUpper(r)
	}, message)

	// Replace vowels in the swapped string
	var result strings.Builder
	for _, r := range swapped {
		if newR, ok := vowelsReplace[r]; ok {
			result.WriteRune(newR)
		} else {
			result.WriteRune(r)
		}
	}

	return result.String()
}