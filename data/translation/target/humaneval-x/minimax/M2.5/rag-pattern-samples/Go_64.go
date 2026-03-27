package main

import (
	"fmt"
	"strings"
)

func VowelsCount(s string) int {
	vowels := "aeiouAEIOU"
	nVowels := 0
	for _, c := range s {
		if strings.ContainsRune(vowels, c) {
			nVowels++
		}
	}
	if len(s) > 0 {
		lastChar := s[len(s)-1]
		if lastChar == 'y' || lastChar == 'Y' {
			nVowels++
		}
	}
	return nVowels
}

func main() {
	// Test cases
	fmt.Println(VowelsCount("abcde")) // Expected: 2
	fmt.Println(VowelsCount("ACEDY")) // Expected: 3
}