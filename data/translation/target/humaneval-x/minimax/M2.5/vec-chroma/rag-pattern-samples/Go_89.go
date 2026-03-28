package main

import (
	"fmt"
	"strings"
)

func Encrypt(s string) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	var result strings.Builder
	
	for _, c := range s {
		idx := strings.Index(alphabet, string(c))
		if idx >= 0 {
			newIdx := (idx + 2*2) % 26 // shift by 4 (2*2)
			result.WriteByte(alphabet[newIdx])
		} else {
			result.WriteRune(c)
		}
	}
	
	return result.String()}

func main() {
	// Test cases
	fmt.Println(Encrypt("hi"))        // expected: lm
	fmt.Println(Encrypt("asdfghjkl")) // expected: ewhjklnop
	fmt.Println(Encrypt("gf"))        // expected: kj
	fmt.Println(Encrypt("et"))        // expected: ix
}