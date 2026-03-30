package main

import (
	"strings"
)

func SelectWords(s string, n int) []string {
	var result []string

	words := strings.Split(s, " ")
	vowels := "aeiouAEIOU"

	for _, word := range words {
		nConsonants := 0
		for _, ch := range word {
			if !strings.ContainsRune(vowels, ch) {
				nConsonants++
			}
		}
		if nConsonants == n {
			result = append(result, word)
		}
	}

	return result
}
