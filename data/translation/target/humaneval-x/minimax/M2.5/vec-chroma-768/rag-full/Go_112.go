package main

import (
	"fmt"
)

func ReverseDelete(s, c string) [2]interface{} {
	// Create a map of characters in c for O(1) lookup
	cMap := make(map[rune]bool)
	for _, char := range c {
		cMap[char] = true
	}

	// Filter out characters from s that are in c
	var result []rune
	for _, char := range s {
		if !cMap[char] {
			result = append(result, char)
		}
	}

	resultStr := string(result)

	// Check if palindrome by comparing with reversed string
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
	result := ReverseDelete("abcde", "ae")
	fmt.Printf("Result: %v\n", result)

	result = ReverseDelete("abcdef", "b")
	fmt.Printf("Result: %v\n", result)

	result = ReverseDelete("abcdedcba", "ab")
	fmt.Printf("Result: %v\n", result)
}
