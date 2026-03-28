package main

import (
	"fmt"
	"sort"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Make a copy to avoid modifying the original slice
	result := make([]string, len(words))
	copy(result, words)

	sort.Slice(result, func(i, j int) bool {
		uniqueI := countUnique(result[i])
		uniqueJ := countUnique(result[j])

		// First compare by number of unique characters (descending)
		if uniqueI != uniqueJ {
			return uniqueI > uniqueJ
		}
		// If equal, compare lexicographically (ascending)
		return result[i] < result[j]
	})

	return result[0]
}

// countUnique returns the number of unique characters in a string
func countUnique(s string) int {
	unique := make(map[rune]struct{})
	for _, r := range s {
		unique[r] = struct{}{}
	}
	return len(unique)
}

func main() {
	// Test cases
	fmt.Println(FindMax([]string{"name", "of", "string"}))      // "string"
	fmt.Println(FindMax([]string{"name", "enam", "game"}))      // "enam"
	fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"}))      // "aaaaaaa"
}