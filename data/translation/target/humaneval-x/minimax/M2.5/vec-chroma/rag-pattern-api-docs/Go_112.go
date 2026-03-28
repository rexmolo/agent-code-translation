package main

import (
	"strings"
)

func ReverseDelete(s, c string) [2]interface{} {
	// Create a map of characters to remove from c
	charToRemove := make(map[rune]bool)
	for _, ch := range c {
		charToRemove[ch] = true
	}

	// Filter characters from s that are not in c
	var result []rune
	for _, ch := range s {
		if !charToRemove[ch] {
			result = append(result, ch)
		}
	}

	resultStr := string(result)

	// Check if palindrome by comparing with reverse
	isPalindrome := resultStr == reverseString(resultStr)

	return [2]interface{}{resultStr, isPalindrome}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Test cases
	result1 := ReverseDelete("abcde", "ae")
	// Expected: ["bcd", false]

	result2 := ReverseDelete("abcdef", "b")
	// Expected: ["acdef", false]

	result3 := ReverseDelete("abcdedcba", "ab")
	// Expected: ["cdedc", true]

	_ = result1
	_ = result2
	_ = result3
}