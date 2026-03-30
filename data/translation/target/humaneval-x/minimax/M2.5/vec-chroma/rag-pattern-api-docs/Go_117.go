package main

import (
	"strings"
)

func SelectWords(s string, n int) []string {
	result := []string{}

	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}

	words := strings.Split(s, " ")

	for _, word := range words {
		nConsonants := 0
		for i := 0; i < len(word); i++ {
			char := strings.ToLower(string(word[i]))
			if !vowels[char] {
				nConsonants++
			}
		}
		if nConsonants == n {
			result = append(result, word)
		}
	}

	return result
}
