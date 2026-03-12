package main

import (
	"fmt"
	"strings"
)

func CountDistinctCharacters(str string) int {
	// Convert string to lowercase to treat characters case-insensitively
	str = strings.ToLower(str)
	
	// Use a map to store distinct characters (Go's equivalent of a set)
	charSet := make(map[rune]struct{})
	
	// Iterate over each character in the string
	for _, r := range str {
		charSet[r] = struct{}{}
	}
	
	return len(charSet)
}

func main() {
	// Test cases
	fmt.Println(CountDistinctCharacters("xyzXYZ")) // Output: 3
	fmt.Println(CountDistinctCharacters("Jerry"))  // Output: 4
}
