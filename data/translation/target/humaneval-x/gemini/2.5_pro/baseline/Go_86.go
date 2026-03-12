package main

import (
	"sort"
	"strings"
)

// AntiShuffle takes a string and returns an ordered version of it.
// An ordered version of a string is one where all words (separated by spaces)
// are replaced by a new word where all characters are arranged in ascending order.
func AntiShuffle(s string) string {
	// Split the input string into words based on spaces.
	words := strings.Split(s, " ")

	// Iterate over the slice of words, sorting each one.
	for i, word := range words {
		// Convert the word to a slice of runes to handle Unicode characters correctly.
		runes := []rune(word)

		// Sort the slice of runes in ascending order based on their code points.
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		// Replace the original word with the sorted version.
		words[i] = string(runes)
	}

	// Join the sorted words back together with spaces.
	return strings.Join(words, " ")
}
