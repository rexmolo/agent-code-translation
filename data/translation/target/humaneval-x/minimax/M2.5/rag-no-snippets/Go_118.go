package main

import (
	"fmt"
)

func GetClosestVowel(word string) string {
	if len(word) < 3 {
		return ""
	}

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// Iterate from len(word)-2 down to 1 (same as Python's range(len(word)-2, 0, -1))
	for i := len(word) - 2; i > 0; i-- {
		current := rune(word[i])
		if vowels[current] {
			// Check if both neighbors are consonants (not vowels)
			if !vowels[rune(word[i+1])] && !vowels[rune(word[i-1])] {
				return string(word[i])
			}
		}
	}
	return ""
}

func main() {
	// Test cases
	fmt.Println(GetClosestVowel("yogurt")) // Expected: "u"
	fmt.Println(GetClosestVowel("FULL"))   // Expected: "U"
	fmt.Println(GetClosestVowel("quick"))  // Expected: ""
	fmt.Println(GetClosestVowel("ab"))    // Expected: ""
}
