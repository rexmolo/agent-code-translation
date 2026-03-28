package main

import (
	"fmt"
	"strconv"
)

// isPalindrome checks if a number is a palindrome by converting to string
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	// Reverse the string
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return s == reversed
}

// EvenOddPalindrome returns a [2]int with the count of even and odd palindromes
// in the range [1, n]
func EvenOddPalindrome(n int) [2]int {
	evenPalindromeCount := 0
	oddPalindromeCount := 0

	for i := 1; i <= n; i++ {
		if !isPalindrome(i) {
			continue
		}
		if i%2 == 1 {
			oddPalindromeCount++
		} else {
			evenPalindromeCount++
		}
	}

	return [2]int{evenPalindromeCount, oddPalindromeCount}
}

func main() {
	// Test examples
	fmt.Println(EvenOddPalindrome(3))  // Output: [1 2]
	fmt.Println(EvenOddPalindrome(12)) // Output: [4 6]
}