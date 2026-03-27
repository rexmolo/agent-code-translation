package main

import (
	"fmt"
	"sort"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Create a sorted copy of the words
	sorted := make([]string, len(words))
	copy(sorted, words)

	sort.Slice(sorted, func(i, j int) bool {
		uqi := countUnique(sorted[i])
		uqj := countUnique(sorted[j])

		// Sort by unique chars descending, then lexicographically ascending
		if uqi != uqj {
			return uqi > uqj
		}
		return sorted[i] < sorted[j]
	})

	return sorted[0]
}

func countUnique(s string) int {
	unique := make(map[rune]bool)
	for _, r := range s {
		unique[r] = true
	}
	return len(unique)
}

func main() {
	// Test cases
	fmt.Println(FindMax([]string{"name", "of", "string"}))   // string
	fmt.Println(FindMax([]string{"name", "enam", "game"}))   // enam
	fmt.Println(FindMax([]string{"aaaaaaa", "bb", "cc"}))   // aaaaaaa
}