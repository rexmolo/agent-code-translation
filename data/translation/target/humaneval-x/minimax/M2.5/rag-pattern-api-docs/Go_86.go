package main

import (
	"fmt"
	"sort"
	"strings"
)

func AntiShuffle(s string) string {
	words := strings.Split(s, " ")

	result := make([]string, len(words))

	for i, word := range words {
		// Convert to rune slice for proper sorting of Unicode characters
		runes := []rune(word)

		// Sort the runes in ascending ASCII order
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		// Convert back to string
		result[i] = string(runes)
	}

	return strings.Join(result, " ")
}

func main() {
	// Test examples
	fmt.Println(AntiShuffle("Hi"))              // Hi
	fmt.Println(AntiShuffle("hello"))           // ehllo
	fmt.Println(AntiShuffle("Hello World!!!"))  // Hello !!!Wdlor
}