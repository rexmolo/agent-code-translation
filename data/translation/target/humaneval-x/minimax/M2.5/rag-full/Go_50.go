package main

import "fmt"

// EncodeShift encodes a string by shifting every character by 5 in the alphabet.
func EncodeShift(s string) string {
	result := make([]byte, len(s))
	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			result[i] = byte(((int(ch)-int('a')+5)%26) + int('a'))
		} else {
			result[i] = byte(ch)
		}
	}
	return string(result)
}

// DecodeShift takes as input string encoded with EncodeShift function. Returns decoded string.
func DecodeShift(s string) string {
	result := make([]byte, len(s))
	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			// Add 26 before modulo to handle negative results
			result[i] = byte(((int(ch)-int('a')-5+26)%26) + int('a'))
		} else {
			result[i] = byte(ch)
		}
	}
	return string(result)
}

func main() {
	// Test the functions
	original := "hello"
	encoded := EncodeShift(original)
	decoded := DecodeShift(encoded)
	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Encoded: %s\n", encoded)
	fmt.Printf("Decoded: %s\n", decoded)
}
