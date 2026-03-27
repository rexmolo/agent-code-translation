package main

import (
	"fmt"
)

func RemoveVowels(text string) string {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	result := make([]rune, 0, len(text))
	for _, r := range text {
		if !vowels[r] {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	// Test cases from docstring
	fmt.Println(RemoveVowels(""))                              // Expected: ''
	fmt.Println(RemoveVowels("abcdef\nghijklm"))             // Expected: 'bcdf\nghjklm'
	fmt.Println(RemoveVowels("abcdef"))                       // Expected: 'bcdf'
	fmt.Println(RemoveVowels("aaaaa"))                        // Expected: ''
	fmt.Println(RemoveVowels("aaBAA"))                        // Expected: 'B'
	fmt.Println(RemoveVowels("zbcd"))                         // Expected: 'zbcd'
}