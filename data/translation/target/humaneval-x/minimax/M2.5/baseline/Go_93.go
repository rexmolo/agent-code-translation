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
		vowelsReplace[v] = v + 2
	}

	// Build result string
	runes := []rune(message)
	for i, r := range runes {
		// Swap case
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else {
			runes[i] = unicode.ToUpper(r)
		}

		// Replace vowel if in map
		if replacement, ok := vowelsReplace[runes[i]]; ok {
			runes[i] = replacement
		}
	}\n
	return string(runes)
}

// The following is for manual testing (not required for the solution)
func main() {
	testCases := []string{"test", "This is a message"}
	for _, tc := range testCases {
		println(Encode(tc))
	}
}
