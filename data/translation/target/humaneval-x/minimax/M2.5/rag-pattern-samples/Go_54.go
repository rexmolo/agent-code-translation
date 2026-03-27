package main

import (
	"fmt"
)

func SameChars(s0 string, s1 string) bool {
	set0 := make(map[rune]bool)
	set1 := make(map[rune]bool)

	for _, r := range s0 {
		set0[r] = true
	}
	for _, r := range s1 {
		set1[r] = true
	}

	if len(set0) != len(set1) {
		return false
	}

	for key := range set0 {
		if !set1[key] {
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
	fmt.Println(SameChars("eabcd", "dddddddabc"))            // false
	fmt.Println(SameChars("abcd", "dddddddabce"))            // false
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}
