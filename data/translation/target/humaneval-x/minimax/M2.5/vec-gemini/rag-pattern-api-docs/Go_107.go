package main

import (
	"fmt"
	"strconv"
)

func EvenOddPalindrome(n int) [2]int {
	isPalindrome := func(num int) bool {
		s := strconv.Itoa(num)
		// Reverse the string
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes) == s
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
	// Test the function
	fmt.Println(EvenOddPalindrome(3))
	fmt.Println(EvenOddPalindrome(12))
}