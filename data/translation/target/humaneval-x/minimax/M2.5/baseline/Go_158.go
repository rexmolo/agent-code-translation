package main

import (
	"sort"
)

func FindMax(words []string) string {
	if len(words) == 0 {
		return ""
	}

	// Helper function to count unique characters
	uniqueChars := func(s string) int {
		seen := make(map[rune]bool)
		for _, c := range s {
			seen[c] = true
		}
		return len(seen)
	}

	// Create a copy to avoid modifying original slice
	result := make([]string, len(words))
	copy(result, words)

	// Sort by: descending unique character count, then ascending lexicographical order
	sort.Slice(result, func(i, j int) bool {
		uniqueI := uniqueChars(result[i])
		uniqueJ := uniqueChars(result[j])
		if uniqueI != uniqueJ {
			return uniqueI > uniqueJ // More unique chars first
		}
		return result[i] < result[j] // Lexicographically smaller first
	})

	return result[0]
}
