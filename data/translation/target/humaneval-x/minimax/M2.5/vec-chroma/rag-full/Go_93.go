package main

import (
	"strings"
)

func Encode(message string) string {
	// Create vowel replacement map
	vowels := "aeiouAEIOU"
	vowelsReplace := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		vowelsReplace[v] = v + 2
	}

	// Swap case of message
	swapped := strings.ToLower(message)
	for i, c := range swapped {
		if c >= 'a' && c <= 'z' {
			swapped = swapped[:i] + string(c-32) + swapped[i+1:]
		}
	}

	// Process each character
	var result []rune
	for _, c := range swapped {
		if replacement, ok := vowelsReplace[c]; ok {
			result = append(result, replacement)
		} else {
			result = append(result, c)
		}
	}

	return string(result)
}