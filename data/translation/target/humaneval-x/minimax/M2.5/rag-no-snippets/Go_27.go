package main

import (
	"strings"
	"unicode"
)

func FlipCase(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsUpper(r) {
			return unicode.ToLower(r)
		}
		if unicode.IsLower(r) {
			return unicode.ToUpper(r)
		}
		return r
	}, str)
}

func main() {
	// Example usage
	result := FlipCase("Hello")
	println(result) // Output: hELLO
}