package main

import (
	"fmt"
	"strconv"
)

func EvenOddPalindrome(n int) [2]int {
	// Helper function to check if a number is a palindrome
	isPalindrome := func(num int) bool {
		s := strconv.Itoa(num)
		// Reverse the string
		reverse := ""
		for i := len(s) - 1; i >= 0; i-- {
			reverse += string(s[i])
		}
		return s == reverse
	}

	evenPalindromeCount := 0
	oddPalindromeCount := 0

	for i := 1; i <= n; i++ {
		if isPalindrome(i) {
			if i%2 == 1 {
				oddPalindromeCount++
			} else {
				evenPalindromeCount++
			}
		}
	}

	return [2]int{evenPalindromeCount, oddPalindromeCount}
}

func main() {
	// Test cases
	fmt.Println(EvenOddPalindrome(3))  // Output: [1 2]
	fmt.Println(EvenOddPalindrome(12)) // Output: [4 6]
}
