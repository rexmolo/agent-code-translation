package main

import (
	"strings"
	"unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")
	if len(parts) == 0 {
		return false
	}

	lastPart := parts[len(parts)-1]
	if len(lastPart) != 1 {
		return false
	}

	lowerChar := unicode.ToLower(rune(lastPart[0]))
	return lowerChar >= 'a' && lowerChar <= 'z'
}

func main() {
	// Test examples
	cases := []string{
		"apple pie",
		"apple pi e",
		"apple pi e ",
		"",
	}
	for _, c := range cases {
		println(CheckIfLastCharIsALetter(c))
	}
}