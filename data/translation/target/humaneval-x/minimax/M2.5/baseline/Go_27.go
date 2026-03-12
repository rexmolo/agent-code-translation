package main

import (
	"fmt"
	"unicode"
)

func FlipCase(str string) string {
	// Convert string to runes to properly handle Unicode characters
	runes := []rune(str)

	// Iterate through each rune and flip its case
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}

	return string(runes)
}

func main() {
	// Test the function
	result := FlipCase("Hello")
	fmt.Println(result)
}