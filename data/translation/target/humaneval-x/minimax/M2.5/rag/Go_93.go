package main

import (
	"strings"
	"unicode"
)

func Encode(message string) string {
	vowels := "aeiouAEIOU"
	vowelsReplace := make(map[rune]rune)

	for _, v := range vowels {
		vowelsReplace[v] = rune(v + 2)
	}

	result := strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		return unicode.ToUpper(r)
	}, message)

	var encoded strings.Builder
	for _, c := range result {
		if newVal, ok := vowelsReplace[c]; ok {
			encoded.WriteRune(newVal)
		} else {
			encoded.WriteRune(c)
		}
	}

	return encoded.String()
}

func main() {
	// Test cases
	// Encode("test") -> "TGST"
	// Encode("This is a message") -> "tHKS KS C MGSSCGG"
}