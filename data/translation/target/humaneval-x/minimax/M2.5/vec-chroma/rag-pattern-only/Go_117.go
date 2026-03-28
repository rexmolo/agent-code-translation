package main

import (
	"fmt"
	"strings"
)

func SelectWords(s string, n int) []string {
	// Use a map for efficient vowel lookup
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	// strings.Fields splits on whitespace (handles multiple spaces)
	words := strings.Fields(s)
	result := make([]string, 0, len(words))

	for _, word := range words {
		nConsonants := 0
		for _, ch := range word {
			// Check if lowercase version is not a vowel
			lowerCh := strings.ToLower(string(ch))
			if _, isVowel := vowels[rune(lowerCh[0])]; !isVowel {
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
	fmt.Println(SelectWords("Mary had a little lamb", 4))  // [little]
	fmt.Println(SelectWords("Mary had a little lamb", 3))   // [Mary lamb]
	fmt.Println(SelectWords("simple white space", 2))      // []
	fmt.Println(SelectWords("Hello world", 4))             // [world]
	fmt.Println(SelectWords("Uncle sam", 3))               // [Uncle]
}