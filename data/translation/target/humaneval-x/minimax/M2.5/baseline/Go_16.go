package main

import (
	"fmt"
	"strings"
)

func CountDistinctCharacters(str string) int {
	lower := strings.ToLower(str)
	seen := make(map[rune]bool)
	for _, r := range lower {
		seen[r] = true
	}
	return len(seen)
}

func main() {
	fmt.Println(CountDistinctCharacters("xyzXYZ")) // Output: 3
	fmt.Println(CountDistinctCharacters("Jerry"))  // Output: 4
}