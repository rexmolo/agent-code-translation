package main

import (
	"fmt"
	"strconv"
	"strings"
)

// isPalindrome checks if an integer is a palindrome by converting
// it to a string and comparing with its reverse
func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	reversed := reverseString(s)
	return s == reversed
}

// reverseString reverses a given string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// EvenOddPalindrome returns an array with the count of even and odd
// integer palindromes in the range [1, n] (inclusive)
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

func main() {
	fmt.Println(EvenOddPalindrome(3))
	fmt.Println(EvenOddPalindrome(12))
}
