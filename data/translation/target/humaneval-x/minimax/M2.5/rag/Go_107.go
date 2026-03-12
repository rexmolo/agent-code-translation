package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func EvenOddPalindrome(n int) [2]int {

evenPalindromeCount := 0
	oddPalindromeCount := 0

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

func main() {
	// Test cases
	result := EvenOddPalindrome(3)
	fmt.Println(result)
	
	result = EvenOddPalindrome(12)
	fmt.Println(result)
}
