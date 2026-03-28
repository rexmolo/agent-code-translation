package main

import "strings"

func RemoveVowels(text string) string {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	var result []string
	for _, r := range text {
		char := string(r)
		if !vowels[strings.ToLower(char)] {
			result = append(result, char)
		}
	}
	return strings.Join(result, "")
}