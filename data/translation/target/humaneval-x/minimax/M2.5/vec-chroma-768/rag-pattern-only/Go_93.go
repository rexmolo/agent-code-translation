package main

import "fmt"

func Encode(message string) string {
	// Define vowels
	vowels := "aeiouAEIOU"

	// Create replacement map: each vowel maps to vowel + 2 in ASCII
	// e becomes g, o becomes q, etc.
	replacements := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		replacements[v] = v + 2
	}

	// Process each character
	result := make([]rune, 0, len(message))
	for _, r := range message {
		var swapped rune
		// Swap case: lowercase->uppercase, uppercase->lowercase
		if r >= 'a' && r <= 'z' {
			swapped = r - 32
		} else if r >= 'A' && r <= 'Z' {
			swapped = r + 32
		} else {
			swapped = r
		}

		// Replace vowel with character 2 places ahead, else keep swapped
		if repl, ok := replacements[swapped]; ok {
			result = append(result, repl)
		} else {
			result = append(result, swapped)
		}
	}

	return string(result)
}

func main() {
	fmt.Println(Encode("test"))
	fmt.Println(Encode("This is a message"))
}
