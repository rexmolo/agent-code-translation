package main

import (
	"sort"
	"strings"
)

func AntiShuffle(s string) string {
	words := strings.Split(s, " ")

	for i, word := range words {
		// Convert to runes to handle characters properly
		runes := []rune(word)

		// Sort the runes by ASCII value
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}
