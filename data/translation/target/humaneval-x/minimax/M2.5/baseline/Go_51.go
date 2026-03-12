package main

import (
	"fmt"
)

func RemoveVowels(text string) string {
	vowels := "aeiouAEIOU"
	result := make([]byte, 0, len(text))
	for i := 0; i < len(text); i++ {
		if !contains(vowels, text[i]) {
			result = append(result, text[i])
		}
	}
	return string(result)
}

func contains(s string, b byte) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == b {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(RemoveVowels(""))
	fmt.Println(RemoveVowels("abcdef\nghijklm"))
	fmt.Println(RemoveVowels("abcdef"))
	fmt.Println(RemoveVowels("aaaaa"))
	fmt.Println(RemoveVowels("aaBAA"))
	fmt.Println(RemoveVowels("zbcd"))
}
