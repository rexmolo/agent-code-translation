package main

import (
	"fmt"
	"strings"
	"unicode"
)

func Solve(s string) string {
	hasLetter := false
	runes := []rune(s)
	
	for i, r := range runes {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				runes[i] = unicode.ToLower(r)
			} else {
				runes[i] = unicode.ToUpper(r)
			}
			hasLetter = true
		}
	}
	
	if !hasLetter {
		// Reverse the string
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
	
	return string(runes)
}

func main() {
	// Test cases
	testCases := []string{"1234", "ab", "#a@C"}
	for _, tc := range testCases {
		fmt.Println(Solve(tc))
	}
}