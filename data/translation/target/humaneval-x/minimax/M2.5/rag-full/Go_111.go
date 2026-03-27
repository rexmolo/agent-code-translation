package main

import (
	"fmt"
	"strings"
)

func Histogram(test string) map[rune]int {
	// Split the string by space
	parts := strings.Split(test, " ")

	// Handle empty string case
	if len(parts) == 0 || (len(parts) == 1 && parts[0] == "") {
		return map[rune]int{}
	}

	// Count occurrences using a map (more efficient than repeated count() calls)
	counts := make(map[rune]int)
	for _, ch := range parts {
		if ch != "" {
			counts[rune(ch[0])]++
		}
	}

	// Find the maximum count
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	// If maxCount is 0 (empty input), return empty map
	if maxCount == 0 {
		return map[rune]int{}
	}

	// Build result with all letters that have max count
	result := make(map[rune]int)
	for ch, count := range counts {
		if count == maxCount {
			result[ch] = count
		}
	}

	return result
}

func main() {
	// Test cases
	testCases := []string{
		"a b c",
		"a b b a",
		"a b c a b",
		"b b b b a",
		"",
	}

	for _, tc := range testCases {
		result := Histogram(tc)
		fmt.Printf("histogram('%s') = %v\n", tc, result)
	}
}