package main

import (
	"fmt"
	"sort"
)

func FindMax(words []string) string {
	// Edge case: empty slice
	if len(words) == 0 {
		return ""
	}

	// Create a copy to avoid modifying the original slice
	result := make([]string, len(words))
	copy(result, words)

	// Sort with custom comparison: first by unique character count (descending),
	// then by lexicographical order (ascending)
	sort.Slice(result, func(i, j int) bool {
		uniqueI := countUnique(result[i])
		uniqueJ := countUnique(result[j])

		// Primary: more unique characters (higher is better)
		if uniqueI != uniqueJ {
			return uniqueI > uniqueJ
		}

		// Secondary: lexicographical order (ascending)
		return result[i] < result[j]
	})

	return result[0]
}

// countUnique returns the number of unique characters in a string
func countUnique(s string) int {
	unique := make(map[rune]struct{})
	for _, c := range s {
		unique[c] = struct{}{}
	}
	return len(unique)
}

func main() {
	// Test cases
	fmt.Println(FindMax([]string{"name", "of", "string"}))   // "string"
	fmt.Println(FindMax([]string{"name", "enam", "game"}))  // "enam"
	fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"}))   // "aaaaaaa"
}