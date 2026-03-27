package main

import (
	"fmt"
	"sort"
	"strings"
)

func AntiShuffle(s string) string {
	words := strings.Split(s, " ")
	
	for i, word := range words {
		// Convert to rune slice for proper sorting (handles Unicode)
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		words[i] = string(runes)
	}
	
	return strings.Join(words, " ")
}

func main() {
	// Test cases
	fmt.Println(AntiShuffle("Hi"))              // Expected: Hi
	fmt.Println(AntiShuffle("hello"))           // Expected: ehllo
	fmt.Println(AntiShuffle("Hello World!!!"))   // Expected: Hello !!!Wdlor
}