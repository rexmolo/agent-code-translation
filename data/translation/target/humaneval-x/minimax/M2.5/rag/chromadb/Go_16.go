package main

import (
	"strings"
)

func CountDistinctCharacters(str string) int {
	lower := strings.ToLower(str)
	distinct := make(map[rune]struct{})
	for _, r := range lower {
		distinct[r] = struct{}{}
	}
	return len(distinct)
}

// For testing
func main() {
	// Test cases
	println(CountDistinctCharacters("xyzXYZ")) // Expected: 3
	println(CountDistinctCharacters("Jerry"))  // Expected: 4
}