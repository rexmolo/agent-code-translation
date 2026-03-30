package main

import "strings"

func CountDistinctCharacters(str string) int {
	lower := strings.ToLower(str)
	chars := map[rune]struct{}{}
	for _, c := range lower {
		chars[c] = struct{}{}
	}
	return len(chars)
}

func main() {
	// Test cases
	println(CountDistinctCharacters("xyzXYZ")) // 3
	println(CountDistinctCharacters("Jerry"))  // 4
}
