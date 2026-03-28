package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	set0 := make(map[rune]bool)
	set1 := make(map[rune]bool)

	// Build set from s0
	for _, r := range s0 {
		set0[r] = true
	}

	// Build set from s1
	for _, r := range s1 {
		set1[r] = true
	}

	// Quick check: if sizes differ, sets can't be equal
	if len(set0) != len(set1) {
		return false
	}

	// Check if all characters in set0 exist in set1
	for k := range set0 {
		if !set1[k] {
			return false
		}
	}

	return true
}

func main() {
	// Test cases from docstring
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc")) // true
	fmt.Println(SameChars("abcd", "dddddddabc"))             // true
	fmt.Println(SameChars("dddddddabc", "abcd"))             // true
	fmt.Println(SameChars("eabcd", "dddddddabc"))           // false
	fmt.Println(SameChars("abcd", "dddddddabce"))           // false
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}