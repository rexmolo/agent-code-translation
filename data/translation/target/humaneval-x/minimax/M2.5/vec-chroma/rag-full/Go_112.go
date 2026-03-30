package main

import (
	"fmt"
	"strings"
)

func ReverseDelete(s, c string) [2]interface{} {
	// Create a map of characters to delete for O(1) lookup
	charMap := make(map[rune]bool)
	for _, char := range c {
		charMap[char] = true
	}

	// Filter: keep only characters not in c
	var result []rune
	for _, char := range s {
		if !charMap[char] {
			result = append(result, char)
		}
	}

	// Convert to string
	resultStr := string(result)

	// Check palindrome: compare with reverse
	isPalindrome := resultStr == reverseString(resultStr)

	return [2]interface{}{resultStr, isPalindrome}
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	// Test cases
	fmt.Println(ReverseDelete("abcde", "ae"))   // [bcd false]
	fmt.Println(ReverseDelete("abcdef", "b"))   // [acdef false]
	fmt.Println(ReverseDelete("abcdedcba", "ab")) // [cdedc true]
}