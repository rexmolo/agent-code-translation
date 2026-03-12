package main

import (
	"fmt"
	"strings"
	"unicode"
)

// SelectWords returns a list of all words from string s that contain exactly
// n consonants, in the order these words appear in the string s.
// If the string s is empty then the function should return an empty list.
// Note: you may assume the input string contains only letters and spaces.
func SelectWords(s string, n int) []string {
	// Handle empty string case explicitly for clarity, though strings.Fields would also work.
	if s == "" {
		return []string{}
	}

	words := strings.Fields(s)
	result := make([]string, 0)

	for _, word := range words {
		consonantCount := 0
		for _, r := range word {
			// Per problem statement, input is only letters and spaces,
			// so any letter that is not a vowel is a consonant.
			lower := unicode.ToLower(r)
			if lower != 'a' && lower != 'e' && lower != 'i' && lower != 'o' && lower != 'u' {
				consonantCount++
			}
		}

		if consonantCount == n {
			result = append(result, word)
		}
	}

	return result
}

// main function to demonstrate the usage of SelectWords
func main() {
	fmt.Printf("%q\n", SelectWords("Mary had a little lamb", 4)) // Expected: ["little"]
	fmt.Printf("%q\n", SelectWords("Mary had a little lamb", 3)) // Expected: ["Mary", "lamb"]
	fmt.Printf("%q\n", SelectWords("simple white space", 2))   // Expected: []
	fmt.Printf("%q\n", SelectWords("Hello world", 4))          // Expected: ["world"]
	fmt.Printf("%q\n", SelectWords("Uncle sam", 3))            // Expected: ["Uncle"]
}
