package main

import "fmt"

// VowelsCount returns the number of vowels in the string.
// Vowels are 'a', 'e', 'i', 'o', 'u' (both cases).
// 'y' is also a vowel, but only when at the end of the word.
func VowelsCount(s string) int {
	isVowel := func(c byte) bool {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}
		return false
	}

	nVowels := 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
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
