package main

import (
	"fmt"
	"strings"
)

func ReverseDelete(s, c string) [2]interface{} {
	// Filter out characters from s that appear in c
	var result strings.Builder
	for _, char := range s {
		if !strings.Contains(c, string(char)) {
			result.WriteString(string(char))
		}
	}
	filtered := result.String()

	// Check if the filtered string is a palindrome
	isPalindrome := true
	runes := []rune(filtered)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			isPalindrome = false
			break
		}
	}

	return [2]interface{}{filtered, isPalindrome}
}

func main() {
	// Test examples
	result1 := ReverseDelete("abcde", "ae")
	fmt.Printf("%v\n", result1)

	result2 := ReverseDelete("abcdef", "b")
	fmt.Printf("%v\n", result2)

	result3 := ReverseDelete("abcdedcba", "ab")
	fmt.Printf("%v\n", result3)
}