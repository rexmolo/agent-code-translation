package main

import "fmt"

func ReverseDelete(s, c string) [2]interface{} {
	// Create a map of characters to delete from c for O(1) lookup
	charsToDelete := make(map[rune]bool)
	for _, char := range c {
		charsToDelete[char] = true
	}

	// Filter s by removing characters that exist in c
	var filteredRunes []rune
	for _, char := range s {
		if !charsToDelete[char] {
			filteredRunes = append(filteredRunes, char)
		}
	}
	result := string(filteredRunes)

	// Check if the result string is a palindrome
	isPalindrome := true
	for i := 0; i < len(result)/2; i++ {
		if result[i] != result[len(result)-1-i] {
			isPalindrome = false
			break
		}
	}

	return [2]interface{}{result, isPalindrome}
}

func main() {
	// Test cases
	result1 := ReverseDelete("abcde", "ae")
	fmt.Printf("Test 1: %v\n", result1)

	result2 := ReverseDelete("abcdef", "b")
	fmt.Printf("Test 2: %v\n", result2)

	result3 := ReverseDelete("abcdedcba", "ab")
	fmt.Printf("Test 3: %v\n", result3)
}