package main

import (
	"strings"
)

func SelectWords(s string, n int) []string {
	if s == "" {
		return []string{}
	}

	words := strings.Split(s, " ")
	var result []string

	for _, word := range words {
		n_consonants := 0
		for _, ch := range word {
			lower := strings.ToLower(string(ch))
			if lower != "a" && lower != "e" && lower != "i" && lower != "o" && lower != "u" {
				n_consonants++
			}
		}
		if n_consonants == n {
			result = append(result, word)
		}
	}

	return result
}
