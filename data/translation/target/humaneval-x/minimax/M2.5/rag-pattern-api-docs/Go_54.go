package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	// Create a set (map) of unique characters from s0
	set0 := make(map[rune]struct{})
	for _, r := range s0 {
		set0[r] = struct{}{}
	}

	// Create a set (map) of unique characters from s1
	set1 := make(map[rune]struct{})
	for _, r := range s1 {
		set1[r] = struct{}{}
	}

	// Quick check: if different number of unique chars, sets can't be equal
	if len(set0) != len(set1) {
		return false
	}

	// Check that all chars in set0 exist in set1
	for r := range set0 {
		if _, ok := set1[r]; !ok {
			return false
		}
	}

	return true
}

func main() {
	// Test cases from docstring
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc")) // true
	fmt.Println(SameChars("abcd", "dddddddabc"))            // true
	fmt.Println(SameChars("dddddddabc", "abcd"))            // true
	fmt.Println(SameChars("eabcd", "dddddddabc"))           // false
	fmt.Println(SameChars("abcd", "dddddddabce"))           // false
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}