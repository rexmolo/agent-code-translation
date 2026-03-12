package main

import (
	"fmt"
	"strings"
)

// CountDistinctCharacters given a string, finds out how many distinct characters (regardless of case) it consists of.
func CountDistinctCharacters(str string) int {
	// A map is used to simulate a set to store unique characters.
	// We use an empty struct as the value to minimize memory usage.
	distinctChars := make(map[rune]struct{})

	// Convert the string to lowercase for case-insensitive comparison.
	lowerStr := strings.ToLower(str)

	// Iterate over the runes (characters) of the lowercase string.
	for _, char := range lowerStr {
		// Add the character to the map. The map will automatically handle duplicates.
		distinctChars[char] = struct{}{}
	}

	// The length of the map gives the count of distinct characters.
	return len(distinctChars)
}

func main() {
	// Example usage based on the Python docstring.
	fmt.Println(CountDistinctCharacters("xyzXYZ")) // Expected output: 3
	fmt.Println(CountDistinctCharacters("Jerry"))  // Expected output: 4
}
