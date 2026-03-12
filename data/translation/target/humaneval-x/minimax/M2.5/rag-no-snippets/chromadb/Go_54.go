package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	// Build a set (map) of unique characters from s0
	set0 := make(map[rune]struct{})
	for _, r := range s0 {
		set0[r] = struct{}{}
	}

	// Build a set (map) of unique characters from s1
	set1 := make(map[rune]struct{})
	for _, r := range s1 {
		set1[r] = struct{}{}
	}

	// Compare the two sets for equality
	if len(set0) != len(set1) {
		return false
	}

	for k := range set0 {
		if _, ok := set1[k]; !ok {
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
	fmt.Println(SameChars("abcd", "dddddddabce"))           // false
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc")) // false
}
