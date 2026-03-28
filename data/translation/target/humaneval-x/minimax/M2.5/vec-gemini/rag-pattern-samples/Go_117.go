package main

import (
	"fmt"
	"strings"
	"unicode"
)

func SelectWords(s string, n int) []string {
	vowels := "aeiou"

	result := []string{}

	for _, word := range strings.Fields(s) {
		nConsonants := 0
		for _, ch := range word {
			if !strings.ContainsRune(vowels, unicode.ToLower(ch)) {
				nConsonants++
			}
		}
		if nConsonants == n {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(SelectWords("Mary had a little lamb", 4)) // [little]
	fmt.Println(SelectWords("Mary had a little lamb", 3)) // [Mary lamb]
	fmt.Println(SelectWords("simple white space", 2))    // []
	fmt.Println(SelectWords("Hello world", 4))           // [world]
	fmt.Println(SelectWords("Uncle sam", 3))              // [Uncle]
}