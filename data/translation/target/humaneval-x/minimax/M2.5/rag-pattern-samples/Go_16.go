package main

import (
	"fmt"
	"strings"
)

func CountDistinctCharacters(str string) int {
	str = strings.ToLower(str)
	uniqueChars := make(map[rune]bool)
	for _, c := range str {
		uniqueChars[c] = true
	}
	return len(uniqueChars)
}

func main() {
	// Test cases
	fmt.Println(CountDistinctCharacters("xyzXYZ")) // Output: 3
	fmt.Println(CountDistinctCharacters("Jerry"))  // Output: 4
}