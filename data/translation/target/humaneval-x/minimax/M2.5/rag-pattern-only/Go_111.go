package main

import (
	"fmt"
	"strings"
)

func Histogram(test string) map[string]int {
	if test == "" {
		return make(map[string]int)
	}

	// Split by space using Fields (handles multiple spaces gracefully)
	list1 := strings.Fields(test)

	if len(list1) == 0 {
		return make(map[string]int)
	}

	// Count occurrences of each letter
	counts := make(map[string]int)
	for _, letter := range list1 {
		counts[letter]++
	}

	// Find maximum count
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	// Build result map with letters that have the maximum count
	result := make(map[string]int)
	for letter, count := range counts {
		if count == maxCount {
			result[letter] = count
		}
	}

	return result
}

func main() {
	// Test cases to verify
	fmt.Println(Histogram("a b c"))       // map[a:1 b:1 c:1]
	fmt.Println(Histogram("a b b a"))     // map[a:2 b:2]
	fmt.Println(Histogram("a b c a b"))   // map[a:2 b:2]
	fmt.Println(Histogram("b b b b a"))   // map[b:4]
	fmt.Println(Histogram(""))           // map[]
}