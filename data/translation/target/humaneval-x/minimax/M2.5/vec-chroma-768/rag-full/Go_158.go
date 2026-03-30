package main

import (
	"fmt"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	maxWord := words[0]
	maxUnique := -1

	for _, word := range words {
		// Count unique characters in the word
		unique := make(map[rune]bool)
		for _, ch := range word {
			unique[ch] = true
		}
		uniqueCount := len(unique)

		// Update max if: more unique chars, or same count with lexicographically smaller word
		if uniqueCount > maxUnique || (uniqueCount == maxUnique && word < maxWord) {
			maxUnique = uniqueCount
			maxWord = word
		}
	}

	return maxWord
}

func main() {
	// Test cases
	fmt.Println(FindMax([]string{"name", "of", "string"})) // "string"
	fmt.Println(FindMax([]string{"name", "enam", "game"})) // "enam"
	fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"}))  // "aaaaaaa"
}