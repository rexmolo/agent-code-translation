package main

import (
	"fmt"
	"strings"
)

// Encrypt takes a string as an argument and returns a string encrypted
// with the alphabet being rotated. The alphabet is rotated in a manner
// such that the letters shift down by four places (2 * 2).
// For example:
//   Encrypt("hi") returns "lm"
//   Encrypt("asdfghjkl") returns "ewhjklnop"
//   Encrypt("gf") returns "kj"
//   Encrypt("et") returns "ix"
func Encrypt(s string) string {
	var result strings.Builder
	shift := 4

	for _, r := range s {
		// The original Python code only encrypts lowercase letters 'a' through 'z'.
		if r >= 'a' && r <= 'z' {
			// Convert rune to a 0-25 index, apply the shift, and wrap around using the modulo operator.
			// Then convert the new index back to a rune.
			shiftedChar := 'a' + (r-'a'+rune(shift))%26
			result.WriteRune(shiftedChar)
		} else {
			// Append any character that is not a lowercase letter unchanged.
			result.WriteRune(r)
		}
	}
	return result.String()
}

// main function to demonstrate the Encrypt function.
func main() {
	fmt.Println(Encrypt("hi"))
	fmt.Println(Encrypt("asdfghjkl"))
	fmt.Println(Encrypt("gf"))
	fmt.Println(Encrypt("et"))
	fmt.Println(Encrypt("Hello World! 123"))
}
