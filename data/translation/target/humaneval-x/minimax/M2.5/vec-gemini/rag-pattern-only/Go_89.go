package main

import (
	"fmt"
)

func Encrypt(s string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var result []rune
	
	for _, c := range s {
		lowerC := toLower(c)
		if idx := findInString(alphabet, lowerC); idx != -1 {
			newIdx := (idx + 4) % 26
			newChar := rune(alphabet[newIdx])
			// Preserve original case
			if c >= 'A' && c <= 'Z' {
				newChar = toUpper(newChar)
			}
			result = append(result, newChar)
		} else {
			result = append(result, c)
		}
	}
	
	return string(result)
}

func toLower(c rune) rune {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

func toUpper(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}
	return c
}

func findInString(s string, c rune) int {
	for i, ch := range s {
		if ch == c {
			return i
		}
	}
	return -1
}

func main() {
	// Test cases
	fmt.Println(Encrypt("hi"))          // Expected: lm
	fmt.Println(Encrypt("asdfghjkl"))   // Expected: ewhjklnop
	fmt.Println(Encrypt("gf"))          // Expected: kj
	fmt.Println(Encrypt("et"))          // Expected: ix
}