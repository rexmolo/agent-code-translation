package main

import (
	"fmt"
	"strings"
)

// Histogram calculates the frequency of space-separated letters and returns a map
// containing only the letter(s) with the highest frequency.
func Histogram(test string) map[rune]int {
	// Use strings.Fields to split the string by any amount of whitespace
	// and discard empty strings, which is more robust than splitting by a single space.
	words := strings.Fields(test)

	// If there are no words after splitting, return an empty map.
	if len(words) == 0 {
		return make(map[rune]int)
	}

	// First pass: count the occurrences of each word (assumed to be a single rune).
	counts := make(map[rune]int)
	for _, word := range words {
		// The Python examples use single-character strings. The Go signature requests
		// a 'rune' key. We take the first rune of the word string to match this.
		runes := []rune(word)
		counts[runes[0]]++
	}

	// Second pass: find the maximum frequency.
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}

	// Third pass: build the result map containing only items with the maximum frequency.
	result := make(map[rune]int)
	for r, count := range counts {
		if count == maxCount {
			result[r] = count
		}
	}

	return result
}

// main is the entry point for the application.
// It demonstrates the Histogram function with examples from the original Python docstring.
func main() {
	fmt.Println(Histogram("a b c"))
	fmt.Println(Histogram("a b b a"))
	fmt.Println(Histogram("a b c a b"))
	fmt.Println(Histogram("b b b b a"))
	fmt.Println(Histogram(""))
}
