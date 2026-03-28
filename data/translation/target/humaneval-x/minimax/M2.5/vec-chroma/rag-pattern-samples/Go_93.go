package main

import (
	"unicode"
)

func Encode(message string) string {
	// Create vowel replacement map
	vowels := "aeiouAEIOU"
	vowelsReplace := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		vowelsReplace[v] = v + 2
	}

	// Swap case and replace vowels
	result := make([]rune, 0, len(message))
	for _, ch := range message {
		// Swap case using unicode
		if unicode.IsLower(ch) {
			ch = unicode.ToUpper(ch)
		} else {
			ch = unicode.ToLower(ch)
		}

		// Replace vowel if exists in map
		if replacement, ok := vowelsReplace[ch]; ok {
			result = append(result, replacement)
		} else {
			result = append(result, ch)
		}
	}

	return string(result)
}

func main() {
	// Test the function
	// fmt.Println(Encode("test"))
	// fmt.Println(Encode("This is a message"))
}