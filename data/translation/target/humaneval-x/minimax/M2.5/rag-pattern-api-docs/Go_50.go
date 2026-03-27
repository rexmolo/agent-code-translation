package main

import (
	"fmt"
)

func DecodeShift(s string) string {
	result := make([]rune, len(s))
	for i, ch := range s {
		// Subtract 5 from the alphabet position, add 26 to handle negative values,
		// then mod 26 and add back 'a' to get the decoded character
		result[i] = rune(((int(ch)-'a'-5+26)%26) + 'a')
	}
	return string(result)
}

func EncodeShift(s string) string {
	result := make([]rune, len(s))
	for i, ch := range s {
		// Add 5 to the alphabet position, mod 26, then add back 'a'
		result[i] = rune(((int(ch)-'a'+5)%26) + 'a')
	}
	return string(result)
}

func main() {
	testStr := "abcdefghijklmnopqrstuvwxyz"
	encoded := EncodeShift(testStr)
	decoded := DecodeShift(encoded)
	fmt.Printf("Original: %s\n", testStr)
	fmt.Printf("Encoded:  %s\n", encoded)
	fmt.Printf("Decoded:  %s\n", decoded)
}
