package main

import (
	"fmt"
)

func ReverseDelete(s, c string) [2]interface{} {
	// Create a map of characters to delete for O(1) lookup
	charMap := make(map[rune]bool)
	for _, char := range c {
		charMap[char] = true
	}

	// Filter out characters that are in c
	var result []rune
	for _, char := range s {
		if !charMap[char] {
			result = append(result, char)
		}
	}

	// Convert result to string
	resultStr := string(result)

	// Check if palindrome
	isPalindrome := true
	n := len(result)
	for i := 0; i < n/2; i++ {
		if result[i] != result[n-1-i] {
			isPalindrome = false
			break
		}
	}

	return [2]interface{}{resultStr, isPalindrome}
}

func main() {
	// Test cases
	fmt.Println(ReverseDelete("abcde", "ae"))     // [bcd false]
	fmt.Println(ReverseDelete("abcdef", "b"))     // [acdef false]
	fmt.Println(ReverseDelete("abcdedcba", "ab")) // [cdedc true]
}
