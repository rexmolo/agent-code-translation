package main

import (
	"fmt"
	"slices"
	"strings"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Sort by: first descending unique character count, then ascending lexicographical order
	slices.SortFunc(words, func(a, b string) int {
		uniqueA := countUniqueChars(a)
		uniqueB := countUniqueChars(b)

		if uniqueA != uniqueB {
			// Return negative for descending order (larger unique count comes first)
			if uniqueA > uniqueB {
				return -1
			}
			return 1
		}
		// If equal unique count, sort lexicographically ascending
		return strings.Compare(a, b)
	})

	return words[0]
}

// countUniqueChars returns the number of unique characters in a string
func countUniqueChars(s string) int {
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