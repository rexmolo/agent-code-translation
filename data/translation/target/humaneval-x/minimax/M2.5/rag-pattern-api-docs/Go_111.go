package main

import (
	"fmt"
	"strings"
)

func Histogram(test string) map[rune]int {
	// Split the string by space to get individual letters
	list1 := strings.Split(test, " ")
	
	// Count occurrences of each letter
	countMap := make(map[rune]int)
	for _, letter := range list1 {
		if letter == "" {
			continue
		}
		// Convert string to rune (single character)
		r := rune(letter[0])
		countMap[r]++
	}
	
	// Find the maximum count
	maxCount := 0
	for _, count := range countMap {
		if count > maxCount {
			maxCount = count
		}
	}
	
	// If no letters (empty input), return empty map
	if maxCount == 0 {
		return make(map[rune]int)
	}
	
	// Build result map with letters that have the maximum count
	result := make(map[rune]int)
	for letter, count := range countMap {
		if count == maxCount {
			result[letter] = count
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
		fmt.Printf("histogram('%s') == %v\n", tc, result)
	}
}
