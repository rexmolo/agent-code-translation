package main

import (
	"fmt"
	"strings"
)

// isPalindrome checks if a number is a palindrome by converting to string
func isPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return s == reversed
}

// EvenOddPalindrome returns the count of even and odd palindromes in range [1, n]
func EvenOddPalindrome(n int) [2]int {
	evenCount := 0
	oddCount := 0

	for i := 1; i <= n; i++ {
		if isPalindrome(i) {
			if i%2 == 0 {
				evenCount++
			} else {
				oddCount++
			}
		}
	}

	return [2]int{evenCount, oddCount}
}

func main() {
	fmt.Println(EvenOddPalindrome(3))  // Output: [1 2]
	fmt.Println(EvenOddPalindrome(12)) // Output: [4 6]
}
