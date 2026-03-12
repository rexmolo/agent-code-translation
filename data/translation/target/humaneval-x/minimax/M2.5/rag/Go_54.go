package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	// Create sets using map[rune]struct{}
	set0 := make(map[rune]struct{})
	set1 := make(map[rune]struct{})

	// Build set from s0
	for _, r := range s0 {
		set0[r] = struct{}{}
	}

	// Build set from s1
	for _, r := range s1 {
		set1[r] = struct{}{}
	}

	// Check if the sets have the same characters
	if len(set0) != len(set1) {
		return false
	}

	for r := range set0 {
		if _, exists := set1[r]; !exists {
			return false
		}
	}

	return true
}

func main() {
	// Test cases
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc")) // true
	fmt.Println(SameChars("abcd", "dddddddabc"))            // true
	fmt.Println(SameChars("dddddddabc", "abcd"))            // true
	fmt.Println(SameChars("eabcd", "dddddddabc"))           // false
	fmt.Println(SameChars("abcd", "dddddddabce"))            // false
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}
