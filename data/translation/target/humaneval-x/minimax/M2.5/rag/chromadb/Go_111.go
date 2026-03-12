package main

import (
	"strings"
)

func Histogram(test string) map[rune]int {
	// Handle empty string case
	if test == "" {
		return map[rune]int{}
	}

	// Split the string by space
	parts := strings.Split(test, " ")

	// Count occurrences of each letter
	counts := make(map[rune]int)
	for _, part := range parts {
		if part == "" {
			continue
		}
		// Convert the part to a rune (lowercase letter)
		r := rune(part[0])
		counts[r]++
	}

	// Find the maximum count
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	// Build result dictionary with letters that have max count
	result := make(map[rune]int)
	for letter, count := range counts {
		if count == maxCount {
			result[letter] = count
		}
	}

	return result
}
