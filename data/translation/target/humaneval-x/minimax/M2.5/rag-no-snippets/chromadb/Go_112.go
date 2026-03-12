package main

import (
	"fmt"
	"strings"
)

func ReverseDelete(s, c string) [2]interface{} {
	// Create a map of characters to remove from c
	charMap := make(map[rune]bool)
	for _, char := range c {
		charMap[char] = true
	}

	// Filter out characters that are in c
	var result strings.Builder
	for _, char := range s {
		if !charMap[char] {
			result.WriteRune(char)
		}
	}

	resultStr := result.String()

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
	fmt.Println(ReverseDelete("abcde", "ae"))
	fmt.Println(ReverseDelete("abcdef", "b"))
	fmt.Println(ReverseDelete("abcdedcba", "ab"))
}
