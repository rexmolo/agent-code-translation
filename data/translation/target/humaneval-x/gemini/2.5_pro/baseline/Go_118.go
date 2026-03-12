package main

import "fmt"

// isVowel checks if a character is a vowel (case-sensitive).
func isVowel(char byte) bool {
	switch char {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// GetClosestVowel finds the closest vowel from the right side of the word that is
// surrounded by two consonants.
func GetClosestVowel(word string) string {
	// A vowel between two consonants requires at least 3 characters.
	if len(word) < 3 {
		return ""
	}

	// Iterate backwards from the second to last character down to the second character.
	// This ensures we are always checking a character with a neighbor on both sides,
	// and we find the "closest from the right" first.
	for i := len(word) - 2; i > 0; i-- {
		// Check for the pattern: consonant-vowel-consonant
		if isVowel(word[i]) && !isVowel(word[i-1]) && !isVowel(word[i+1]) {
			return string(word[i])
		}
	}

	// If the loop finishes, no such vowel was found.
	return ""
}

// main function to demonstrate the usage of GetClosestVowel with examples.
func main() {
	fmt.Println(GetClosestVowel("yogurt"))
	fmt.Println(GetClosestVowel("FULL"))
	fmt.Println(GetClosestVowel("quick"))
	fmt.Println(GetClosestVowel("ab"))
	fmt.Println(GetClosestVowel("rhythm"))
}
