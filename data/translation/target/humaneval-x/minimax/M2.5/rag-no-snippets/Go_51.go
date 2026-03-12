package main

import (
	"strings"
)

func RemoveVowels(text string) string {
	vowels := "aeiou"

	var builder strings.Builder
	for _, ch := range text {
		if !strings.Contains(vowels, strings.ToLower(string(ch))) {
			builder.WriteRune(ch)
		}
	}
	return builder.String()
}
