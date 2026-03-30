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

	// Iterate from right side (len(word)-2) down to 1
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
	fmt.Println(GetClosestVowel("yogurt")) // Output: u
	fmt.Println(GetClosestVowel("FULL"))   // Output: U
	fmt.Println(GetClosestVowel("quick"))  // Output: 
	fmt.Println(GetClosestVowel("ab"))     // Output: 
}
