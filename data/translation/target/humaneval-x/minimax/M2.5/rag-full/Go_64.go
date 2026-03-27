package main

import "fmt"

func VowelsCount(s string) int {
	// Count vowels (a, e, i, o, u) - case insensitive
	// 'y' is also a vowel, but only when it appears at the end
	nVowels := 0
	for _, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
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
	// Test cases from the docstring
	fmt.Println(VowelsCount("abcde")) // Expected: 2
	fmt.Println(VowelsCount("ACEDY")) // Expected: 3
}