package main

import "fmt"

func SameChars(s0 string, s1 string) bool {
	set0 := make(map[rune]bool)
	set1 := make(map[rune]bool)

	for _, c := range s0 {
		set0[c] = true
	}
	for _, c := range s1 {
		set1[c] = true
	}

	if len(set0) != len(set1) {
		return false
	}

	for c := range set0 {
		if !set1[c] {
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
