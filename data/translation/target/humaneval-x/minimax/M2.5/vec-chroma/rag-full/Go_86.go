package main

import (
	"sort"
	"strings"
)

func AntiShuffle(s string) string {
	words := strings.Split(s, " ")
	sortedWords := make([]string, len(words))

	for i, word := range words {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedWords[i] = string(runes)
	}

	return strings.Join(sortedWords, " ")
}