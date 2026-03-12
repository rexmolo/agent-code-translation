package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	// Create maps to store unique characters from each string
	chars0 := make(map[rune]bool)
	chars1 := make(map[rune]bool)

	// Populate chars0 with unique characters from s0
	for _, r := range s0 {
		chars0[r] = true
	}

	// Populate chars1 with unique characters from s1
	for _, r := range s1 {
		chars1[r] = true
	}

	// If different number of unique characters, they can't have the same set
	if len(chars0) != len(chars1) {
		return false
	}

	// Check if all characters in chars0 exist in chars1
	for r := range chars0 {
		if !chars1[r] {
			return false
		}
	}

	return true
}

func main() {
	// Test cases from docstring
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc")) // true
	fmt.Println(SameChars("abcd", "dddddddabc"))              // true
	fmt.Println(SameChars("dddddddabc", "abcd"))              // true
	fmt.Println(SameChars("eabcd", "dddddddabc"))             // false
	fmt.Println(SameChars("abcd", "dddddddabce"))             // false
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}
