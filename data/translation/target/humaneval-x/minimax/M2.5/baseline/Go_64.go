package main

import (
	"fmt"
	"strings"
)

func VowelsCount(s string) int {
	vowels := "aeiouAEIOU"
	nVowels := 0

	for _, c := range s {
		if strings.ContainsRune(vowels, rune(c)) {
			nVowels++
		}
	}

	// Check if last character is 'y' or 'Y' - only counts as vowel at the end
	if len(s) > 0 {
		lastChar := s[len(s)-1]
		if lastChar == 'y' || lastChar == 'Y' {
			nVowels++
		}
	}

	return nVowels
}

func main() {
	fmt.Println(VowelsCount("abcde")) // Output: 2
	fmt.Println(VowelsCount("ACEDY")) // Output: 3
	fmt.Println(VowelsCount(""))      // Output: 0
	fmt.Println(VowelsCount("y"))     // Output: 1
}
