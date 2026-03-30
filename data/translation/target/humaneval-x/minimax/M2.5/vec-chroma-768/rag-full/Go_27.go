package main

import "strings"

func FlipCase(str string) string {
	// Convert to rune slice to properly handle Unicode characters
	runes := []rune(str)
	for i, r := range runes {
		switch {
		case r >= 'A' && r <= 'Z':
			// Uppercase letter: convert to lowercase
			runes[i] = r + 32
		case r >= 'a' && r <= 'z':
			// Lowercase letter: convert to uppercase
			runes[i] = r - 32
		default:
			// Non-letter characters remain unchanged
		}
	}
	return string(runes)
}

// The following is a test function to verify the translation
func main() {
	// Simple test - can be run with: go run main.go
	// Or import and test FlipCase directly
	testCases := []string{"Hello", "World", "PyThOn", "123!@#"}
	for _, tc := range testCases {
		println(FlipCase(tc))
	}
}