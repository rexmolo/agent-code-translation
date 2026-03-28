package main

import "strings"

// VowelsCount takes a string representing a word as input and returns
// the number of vowels in the string.
// Vowels in this case are 'a', 'e', 'i', 'o', 'u'. Here, 'y' is also a
// vowel, but only when it is at the end of the given word.
func VowelsCount(s string) int {
	vowels := "aeiouAEIOU"
	nVowels := 0

	for _, c := range s {
		if strings.ContainsRune(vowels, c) {
			nVowels++
		}
	}

	// Check if the last character is 'y' or 'Y'
	if len(s) > 0 && (s[len(s)-1] == 'y' || s[len(s)-1] == 'Y') {
		nVowels++
	}

	return nVowels
}
