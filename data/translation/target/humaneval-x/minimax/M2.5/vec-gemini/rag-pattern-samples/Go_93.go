package main

import (
	"strings"
	"unicode"
)

func Encode(message string) string {
	vowels := "aeiouAEIOU"
	vowelsReplace := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		vowelsReplace[v] = v + 2
	}

	message = swapCase(message)

	var result strings.Builder
	result.Grow(len(message))

	for _, ch := range message {
		if replacement, ok := vowelsReplace[ch]; ok {
			result.WriteRune(replacement)
		} else {
			result.WriteRune(ch)
		}
	}

	return result.String()
}

func swapCase(s string) string {
	var result strings.Builder
	result.Grow(len(s))
	for _, r := range s {
		if unicode.IsUpper(r) {
			result.WriteRune(unicode.ToLower(r))
		} else if unicode.IsLower(r) {
			result.WriteRune(unicode.ToUpper(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}

func main() {
	// Test cases
	println(Encode("test"))
	println(Encode("This is a message"))
}
