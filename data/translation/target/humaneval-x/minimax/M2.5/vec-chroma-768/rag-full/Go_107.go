package main

import (
	"strconv"
)

// isPalindrome checks if a number is a palindrome by converting it to a string
// and comparing it with its reverse.
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	// Reverse the string
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return s == reversed
}

func EvenOddPalindrome(n int) [2]int {

evenPalindromeCount := 0
	oddPalindromeCount := 0

	for i := 1; i <= n; i++ {
		if i%2 == 1 && isPalindrome(i) {
			oddPalindromeCount++
		} else if i%2 == 0 && isPalindrome(i) {
			evenPalindromeCount++
		}
	}

	return [2]int{evenPalindromeCount, oddPalindromeCount}
}