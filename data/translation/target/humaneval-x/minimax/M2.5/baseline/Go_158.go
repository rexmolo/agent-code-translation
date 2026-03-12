package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Create a copy to avoid modifying the original slice
	wordsCopy := make([]string, len(words))
	copy(wordsCopy, words)

	// Sort with custom comparator:
	// - First by descending number of unique characters
	// - Then by ascending lexicographical order
	sort.Slice(wordsCopy, func(i, j int) bool {
		uniqueI := countUnique(wordsCopy[i])
		uniqueJ := countUnique(wordsCopy[j])

		if uniqueI != uniqueJ {
			return uniqueI > uniqueJ // More unique characters first
		}
		return strings.Compare(wordsCopy[i], wordsCopy[j]) < 0 // Lexicographical order
	})

	return wordsCopy[0]
}

func countUnique(s string) int {
	unique := make(map[rune]bool)
	for _, r := range s {
		unique[r] = true
	}
	return len(unique)
}

func main() {
	fmt.Println(FindMax([]string{"name", "of", "string"}))   // string
	fmt.Println(FindMax([]string{"name", "enam", "game"}))   // enam
	fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"}))   // aaaaaaa
}