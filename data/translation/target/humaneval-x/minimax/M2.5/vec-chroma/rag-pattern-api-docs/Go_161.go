package main

import (
	"unicode"
)

func Solve(s string) string {
	runes := []rune(s)
	hasLetter := false
	
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
		// Reverse the string in place
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}
	
	return string(runes)
}

func main() {
	// Test cases
	println(Solve("1234")) // "4321"
	println(Solve("ab"))   // "AB"
	println(Solve("#a@C")) // "#A@c"
}