package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	reverse := ""
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return s == reverse
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
	fmt.Println(EvenOddPalindrome(3))
	fmt.Println(EvenOddPalindrome(12))
}
