package main

import (
	"fmt"
	"strings"
)

// ReverseDelete removes characters from s that are in c and checks if the result is a palindrome.
func ReverseDelete(s, c string) [2]interface{} {
	// Create a set of characters from c for efficient lookup.
	// Using struct{} is more memory-efficient than bool for a set.
	charsToDelete := make(map[rune]struct{}, len(c))
	for _, char := range c {
		charsToDelete[char] = struct{}{}
	}

	// Use a strings.Builder to efficiently construct the new string.
	var builder strings.Builder
	for _, char := range s {
		if _, found := charsToDelete[char]; !found {
			builder.WriteRune(char)
		}
	}
	filteredS := builder.String()

	// Check if the resulting string is a palindrome.
	isPalindrome := true
	runes := []rune(filteredS)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			isPalindrome = false
			break
		}
	}

	// Return the result string and the palindrome check as a fixed-size array of interfaces.
	return [2]interface{}{filteredS, isPalindrome}
}

func main() {
	// Example 1
	result1 := ReverseDelete("abcde", "ae")
	fmt.Printf("Result: ('%s', %v)\n", result1[0], result1[1])

	// Example 2
	result2 := ReverseDelete("abcdef", "b")
	fmt.Printf("Result: ('%s', %v)\n", result2[0], result2[1])

	// Example 3
	result3 := ReverseDelete("abcdedcba", "ab")
	fmt.Printf("Result: ('%s', %v)\n", result3[0], result3[1])
}
