package main

import (
	"fmt"
	"strings"
)

func RemoveVowels(text string) string {
	vowelsLower := "aeiou"
	vowelsUpper := "AEIOU"

	var result strings.Builder
	for _, r := range text {
		if !strings.ContainsRune(vowelsLower, r) && !strings.ContainsRune(vowelsUpper, r) {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func main() {
	fmt.Println(RemoveVowels(""))
	fmt.Println(RemoveVowels("abcdef\nghijklm"))
	fmt.Println(RemoveVowels("abcdef"))
	fmt.Println(RemoveVowels("aaaaa"))
	fmt.Println(RemoveVowels("aaBAA"))
	fmt.Println(RemoveVowels("zbcd"))
}
