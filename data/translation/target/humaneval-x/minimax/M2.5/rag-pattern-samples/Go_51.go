package main

import (
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
		lower := unicode.ToLower(r)
		if !vowels[lower] {
			result = append(result, r)
		}
	}
	return string(result)
}