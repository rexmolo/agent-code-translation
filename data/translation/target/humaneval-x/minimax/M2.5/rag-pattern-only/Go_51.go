package main

import "fmt"

func RemoveVowels(text string) string {
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	var result []rune
	for _, r := range text {
		if !vowels[r] {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	// Test cases from the docstring
	fmt.Println(RemoveVowels(""))
	fmt.Println(RemoveVowels("abcdef\nghijklm"))
	fmt.Println(RemoveVowels("abcdef"))
	fmt.Println(RemoveVowels("aaaaa"))
	fmt.Println(RemoveVowels("aaBAA"))
	fmt.Println(RemoveVowels("zbcd"))
}