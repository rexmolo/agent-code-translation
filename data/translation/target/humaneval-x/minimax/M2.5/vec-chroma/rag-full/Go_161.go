package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Solve(s string) string {
	// Convert string to rune slice for proper Unicode handling
	runes := []rune(s)
	hasLetter := false

	for i, r := range runes {
		if unicode.IsLetter(r) {
			hasLetter = true
			if unicode.IsUpper(r) {
				runes[i] = []rune(strings.ToLower(string(r)))[0]
			} else {
				runes[i] = []rune(strings.ToUpper(string(r)))[0]
			}
		}
	}

	// If no letters found, reverse the string
	if !hasLetter {
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}

	return string(runes)
}

func main() {
	// Test cases
	fmt.Println(Solve("1234")) // Expected: "4321"
	fmt.Println(Solve("ab"))   // Expected: "AB"
	fmt.Println(Solve("#a@C")) // Expected: "#A@c"
}
