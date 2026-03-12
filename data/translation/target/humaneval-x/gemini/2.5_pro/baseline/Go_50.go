package main

import (
	"fmt"
)

// DecodeShift takes a string encoded with a Caesar cipher (shift of 5)
// and returns the decoded string. It assumes the input string consists
// of lowercase English alphabet characters.
func DecodeShift(s string) string {
	// Convert the string to a slice of runes for easier character manipulation.
	runes := []rune(s)
	for i, r := range runes {
		// Calculate the 0-based position in the alphabet (a=0, b=1, ...).
		pos := int(r - 'a')

		// Shift the position back by 5. We add 26 before the modulo
		// to handle negative results correctly, ensuring the result
		// wraps around the 26-letter alphabet.
		// Go's % operator can return a negative result for negative inputs,
		// unlike Python's. (pos - 5 + 26) % 26 is an idiomatic way
		// to get a true mathematical modulus.
		newPos := (pos - 5 + 26) % 26

		// Convert the new position back to a rune and update the slice.
		runes[i] = rune('a' + newPos)
	}
	// Convert the slice of runes back to a string.
	return string(runes)
}

func main() {
	// Example usage to demonstrate the function
	encodedString1 := "fghij"
	decodedString1 := DecodeShift(encodedString1)
	fmt.Printf("Encoded: %s -> Decoded: %s\n", encodedString1, decodedString1)

	encodedString2 := "mjqqt"
	decodedString2 := DecodeShift(encodedString2)
	fmt.Printf("Encoded: %s -> Decoded: %s\n", encodedString2, decodedString2)

	// Example that involves wrapping around the beginning of the alphabet
	encodedString3 := "abcde"
	decodedString3 := DecodeShift(encodedString3)
	fmt.Printf("Encoded: %s -> Decoded: %s\n", encodedString3, decodedString3)
}
