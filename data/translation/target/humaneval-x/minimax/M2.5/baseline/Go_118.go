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

	for i := len(word) - 2; i > 0; i-- {
		current := rune(word[i])
		if vowels[current] {
			prev := rune(word[i-1])
			next := rune(word[i+1])
			if !vowels[prev] && !vowels[next] {
				return string(current)
			}
		}
	}

	return ""
}

func main() {
	// Test cases
	fmt.Println(GetClosestVowel("yogurt")) // Output: u
	fmt.Println(GetClosestVowel("FULL"))   // Output: U
	fmt.Println(GetClosestVowel("quick"))   // Output: 
	fmt.Println(GetClosestVowel("ab"))     // Output: 
}