package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	s0Chars := make(map[rune]struct{})
	s1Chars := make(map[rune]struct{})

	for _, r := range s0 {
		s0Chars[r] = struct{}{}
	}
	for _, r := range s1 {
		s1Chars[r] = struct{}{}
	}

	if len(s0Chars) != len(s1Chars) {
		return false
	}

	for k := range s0Chars {
		if _, ok := s1Chars[k]; !ok {
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