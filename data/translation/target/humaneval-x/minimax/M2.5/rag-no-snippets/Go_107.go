package main

import (
	"strconv"
)

func EvenOddPalindrome(n int) [2]int {
	// Helper function to check if a number is a palindrome
	isPalindrome := func(num int) bool {
		s := strconv.Itoa(num)
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			if runes[i] != runes[j] {
				return false
			}
		}
		return true
	}

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
