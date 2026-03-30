package main

import (
	"fmt"
)

// VowelsCount takes a string and returns the number of vowels in it.
// Vowels are 'a', 'e', 'i', 'o', 'u' (case-insensitive).
// 'y' or 'Y' is also counted as a vowel, but only when it appears
// at the end of the word.
func VowelsCount(s string) int {
	vowelMap := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	nVowels := 0

	// Count regular vowels
	for _, c := range s {
		if vowelMap[c] {
			nVowels++
		}
	}

	// Check if last character is 'y' or 'Y'
	if len(s) > 0 && (s[len(s)-1] == 'y' || s[len(s)-1] == 'Y') {
		nVowels++
	}

	return nVowels
}

func main() {
	fmt.Println(VowelsCount("abcde")) // Output: 2
	fmt.Println(VowelsCount("ACEDY")) // Output: 3
}
