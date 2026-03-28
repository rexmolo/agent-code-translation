package main

import (
	"strings"
)

func ReverseDelete(s, c string) [2]interface{} {
	// Filter characters: remove all characters in c from s
	var result []rune
	for _, char := range s {
		if !strings.Contains(c, string(char)) {
			result = append(result, char)
		}
	}

	// Convert to string
	resultStr := string(result)

	// Check palindrome by comparing string with its reverse
	isPalindrome := true
	for i := 0; i < len(resultStr)/2; i++ {
		if resultStr[i] != resultStr[len(resultStr)-1-i] {
			isPalindrome = false
			break
		}
	}

	return [2]interface{}{resultStr, isPalindrome}
}