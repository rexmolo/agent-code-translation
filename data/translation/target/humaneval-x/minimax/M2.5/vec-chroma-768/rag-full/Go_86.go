package main

import (
	"fmt"
	"sort"
	"strings"
)

func AntiShuffle(s string) string {
	words := strings.Split(s, " ")

	for i, word := range words {
		// Convert word to rune slice for proper Unicode handling
		runes := []rune(word)

		// Sort runes by ASCII value
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		// Convert back to string
		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func main() {
	// Test cases
	fmt.Println(AntiShuffle("Hi"))
	fmt.Println(AntiShuffle("hello"))
	fmt.Println(AntiShuffle("Hello World!!!"))
}
