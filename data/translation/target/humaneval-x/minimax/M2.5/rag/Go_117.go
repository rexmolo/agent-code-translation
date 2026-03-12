package main

import (
	"strings"
)

func SelectWords(s string, n int) []string {
	// If the string is empty, return empty list
	if s == "" {
		return []string{}
	}

	result := []string{}
	words := strings.Split(s, " ")

	for _, word := range words {
		nConsonants := 0
		for _, ch := range word {
			lower := strings.ToLower(string(ch))
			if lower != "a" && lower != "e" && lower != "i" && lower != "o" && lower != "u" {
				nConsonants++
			}
		}
		if nConsonants == n {
			result = append(result, word)
		}
	}

	return result
}