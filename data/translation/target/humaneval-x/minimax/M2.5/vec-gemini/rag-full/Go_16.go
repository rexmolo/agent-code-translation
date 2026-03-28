package main

import (
	"fmt"
	"strings"
)

func CountDistinctCharacters(str string) int {
	// Convert string to lowercase
	lower := strings.ToLower(str)
	
	// Use map to track unique characters (Go's equivalent of Python's set)
	unique := map[rune]struct{}{}
	
	// Iterate over each character (rune) in the string
	for _, c := range lower {
		unique[c] = struct{}{}
	}
	
	// Return the count of distinct characters
	return len(unique)
}

func main() {
	// Test the function
	fmt.Println(CountDistinctCharacters("xyzXYZ")) // Output: 3
	fmt.Println(CountDistinctCharacters("Jerry"))  // Output: 4
}
