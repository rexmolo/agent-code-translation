package main

import "fmt"

func Encrypt(s string) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	result := make([]byte, 0, len(s))
	
	for _, c := range s {
		// Check if lowercase letter (a-z)
		if c >= 'a' && c <= 'z' {
			// Find index and shift by 4 (2*2)
			idx := c - 'a'
			newIdx := (idx + 4) % 26
			result = append(result, 'a'+byte(newIdx))
		} else {
			// Keep non-lowercase letters as-is (including uppercase)
			result = append(result, byte(c))
		}
	}
	
	return string(result)
}

func main() {
	fmt.Println(Encrypt("hi"))        // Expected: lm
	fmt.Println(Encrypt("asdfghjkl")) // Expected: ewhjklnop
	fmt.Println(Encrypt("gf"))        // Expected: kj
	fmt.Println(Encrypt("et"))        // Expected: ix
}
