package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	set0 := make(map[rune]struct{})
	for _, r := range s0 {
		set0[r] = struct{}{}
	}

	set1 := make(map[rune]struct{})
	for _, r := range s1 {
		set1[r] = struct{}{}
	}

	// Quick check: if sizes differ, sets are different
	if len(set0) != len(set1) {
		return false
	}

	// Check if all characters in set0 exist in set1
	for r := range set0 {
		if _, exists := set1[r]; !exists {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc")) // true
	fmt.Println(SameChars("abcd", "dddddddabc"))             // true
	fmt.Println(SameChars("dddddddabc", "abcd"))             // true
	fmt.Println(SameChars("eabcd", "dddddddabc"))            // false
	fmt.Println(SameChars("abcd", "dddddddabce"))            // false
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}
