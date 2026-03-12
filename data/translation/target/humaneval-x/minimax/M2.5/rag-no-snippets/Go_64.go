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

	// Check if string is not empty and ends with 'y' or 'Y'
	if len(s) > 0 {
		lastChar := s[len(s)-1]
		if lastChar == 'y' || lastChar == 'Y' {
			nVowels++
		}
	}

	return nVowels
}

func main() {
	fmt.Println(VowelsCount("abcde"))  // 2
	fmt.Println(VowelsCount("ACEDY"))  // 3
}
