package main

import (
	"fmt"
	"strings"
)

func CountDistinctCharacters(str string) int {
	lower := strings.ToLower(str)
	uniqueChars := make(map[rune]struct{})
	for _, r := range lower {
		uniqueChars[r] = struct{}{}
	}
	return len(uniqueChars)
}

func main() {
	// Test examples
	fmt.Println(CountDistinctCharacters("xyzXYZ")) // Output: 3
	fmt.Println(CountDistinctCharacters("Jerry"))  // Output: 4
}