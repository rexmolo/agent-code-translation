package main

import (
	"fmt"
	"sort"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Create a copy to avoid modifying the original slice
	wordsCopy := make([]string, len(words))
	copy(wordsCopy, words)

	sort.Slice(wordsCopy, func(i, j int) bool {
		// Count unique characters in wordsCopy[i]
		uniqueI := make(map[rune]bool)
		for _, c := range wordsCopy[i] {
			uniqueI[c] = true
		}
		lenI := len(uniqueI)

		// Count unique characters in wordsCopy[j]
		uniqueJ := make(map[rune]bool)
		for _, c := range wordsCopy[j] {
			uniqueJ[c] = true
		}
		lenJ := len(uniqueJ)

		// Sort by: more unique characters first (descending),
		// then lexicographically (ascending) for ties
		if lenI != lenJ {
			return lenI > lenJ
		}
		return wordsCopy[i] < wordsCopy[j]
	})

	return wordsCopy[0]
}

func main() {
	// Test cases
	fmt.Println(FindMax([]string{"name", "of", "string"})) // "string"
	fmt.Println(FindMax([]string{"name", "enam", "game"})) // "enam"
	fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"})) // "aaaaaaa"
}