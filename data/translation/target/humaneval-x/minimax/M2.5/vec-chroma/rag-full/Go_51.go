package main

import (
	"fmt"
	"strings"
	"unicode"
)

func RemoveVowels(text string) string {
	var result []rune
	vowels := "aeiou"
	for _, r := range text {
		if !strings.ContainsRune(vowels, unicode.ToLower(r)) {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	var text string
	fmt.Scan(&text)
	fmt.Print(RemoveVowels(text))
}
