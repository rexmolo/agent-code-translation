package main

import (
	"strings"
	"unicode"
)

func RemoveVowels(text string) string {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	var result []rune
	for _, r := range text {
		if !vowels[unicode.ToLower(r)] {
			result = append(result, r)
		}
	}
	return string(result)
}
