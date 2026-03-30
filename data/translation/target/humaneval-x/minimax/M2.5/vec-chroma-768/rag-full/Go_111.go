package main

import "strings"

func Histogram(test string) map[rune]int {
	if test == "" {
		return map[rune]int{}
	}

	// Split by space to get individual letters
	words := strings.Split(test, " ")

	// Count frequency of each letter
	count := make(map[rune]int)
	for _, word := range words {
		if word != "" {
			count[rune(word[0])]++
		}
	}

	// Find maximum count
	maxCount := 0
	for _, v := range count {
		if v > maxCount {
			maxCount = v
		}
	}

	// Build result with only letters that have max count
	result := make(map[rune]int)
	for k, v := range count {
		if v == maxCount {
			result[k] = v
		}
	}

	return result
}

func main() {
	// Test cases can be added here for verification
}
