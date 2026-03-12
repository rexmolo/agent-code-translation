package main

import (
	"fmt"
	"strconv"
)

// isPalindrome checks if an integer is a palindrome.
// It works by converting the integer to a string and comparing it with its reverse.
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	runes := []rune(s)
	// Reverse the rune slice in-place
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Compare the original string with the reversed string
	return s == string(runes)
}

// EvenOddPalindrome counts the number of even and odd integer palindromes
// that fall within the range [1, n], inclusive.
// It returns a 2-element array where the first element is the count of
// even palindromes and the second is the count of odd palindromes.
func EvenOddPalindrome(n int) [2]int {
	var evenPalindromeCount int
	var oddPalindromeCount int

	for i := 1; i <= n; i++ {
		if isPalindrome(i) {
			if i%2 == 0 {
				evenPalindromeCount++
			} else {
				oddPalindromeCount++
			}
		}
	}

	return [2]int{evenPalindromeCount, oddPalindromeCount}
}

// main function to run and test the code
func main() {
	fmt.Println("--- Example 1 ---")
	result1 := EvenOddPalindrome(3)
	fmt.Printf("Input: 3\nOutput: %v\n", result1) // Expected: [1 2]

	fmt.Println("\n--- Example 2 ---")
	result2 := EvenOddPalindrome(12)
	fmt.Printf("Input: 12\nOutput: %v\n", result2) // Expected: [4 6]

	fmt.Println("\n--- Test with a larger number ---")
	result3 := EvenOddPalindrome(150)
	fmt.Printf("Input: 150\nOutput: %v\n", result3) // Palindromes: 1-9, 11, 22, 33, 44, 55, 66, 77, 88, 99, 101, 111, 121, 131, 141
}
