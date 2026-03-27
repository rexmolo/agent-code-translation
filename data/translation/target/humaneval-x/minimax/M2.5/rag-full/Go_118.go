package main

import "fmt"

func GetClosestVowel(word string) string {
	if len(word) < 3 {
		return ""
	}

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	// Iterate from right side (excluding last two chars) going left to position 1
	for i := len(word) - 2; i > 0; i-- {
		char := rune(word[i])
		if vowels[char] {
			leftChar := rune(word[i-1])
			rightChar := rune(word[i+1])
			if !vowels[leftChar] && !vowels[rightChar] {
				return string(char)
			}
		}
	}

	return ""
}

func main() {
	// Test cases
	fmt.Println(GetClosestVowel("yogurt"))  // Expected: "u"
	fmt.Println(GetClosestVowel("FULL"))   // Expected: "U"
	fmt.Println(GetClosestVowel("quick"))  // Expected: ""
	fmt.Println(GetClosestVowel("ab"))     // Expected: ""
}
