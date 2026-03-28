package main

import "fmt"

func ReverseDelete(s, c string) [2]interface{} {
	// Build a map for efficient lookup of characters to delete
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

	// Check if palindrome
	str := string(result)
	isPalindrome := true
	n := len(result)
	for i := 0; i < n/2; i++ {
		if result[i] != result[n-1-i] {
			isPalindrome = false
			break
		}
	}

	return [2]interface{}{str, isPalindrome}
}

func main() {
	// Test cases
	result1 := ReverseDelete("abcde", "ae")
	fmt.Println(result1) // [bcd false]

	result2 := ReverseDelete("abcdef", "b")
	fmt.Println(result2) // [acdef false]

	result3 := ReverseDelete("abcdedcba", "ab")
	fmt.Println(result3) // [cdedc true]
}