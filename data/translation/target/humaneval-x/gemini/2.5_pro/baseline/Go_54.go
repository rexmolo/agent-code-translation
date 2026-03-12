package main

import "fmt"

// SameChars checks if two words have the same unique characters.
func SameChars(s0 string, s1 string) bool {
	set0 := make(map[rune]struct{})
	for _, r := range s0 {
		set0[r] = struct{}{}
	}

	set1 := make(map[rune]struct{})
	for _, r := range s1 {
		set1[r] = struct{}{}
	}

	if len(set0) != len(set1) {
		return false
	}

	for r := range set0 {
		if _, ok := set1[r]; !ok {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddeddabc"))
	fmt.Println(SameChars("abcd", "dddddddabc"))
	fmt.Println(SameChars("dddddddabc", "abcd"))
	fmt.Println(SameChars("eabcd", "dddddddabc"))
	fmt.Println(SameChars("abcd", "dddddddabce"))
	fmt.Println(SameChars("eabcdzzzz", "dddzzzzzzzddddabc"))
}
