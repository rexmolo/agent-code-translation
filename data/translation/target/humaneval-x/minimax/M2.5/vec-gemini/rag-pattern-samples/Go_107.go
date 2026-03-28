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
	evenCount := 0
	oddCount := 0

	for i := 1; i <= n; i++ {
		if isPalindrome(i) {
			if i%2 == 1 {
				oddCount++
			} else {
				evenCount++
			}
		}
	}
	return [2]int{evenCount, oddCount}
}

func main() {
	// Test cases from the problem
	fmt.Println(EvenOddPalindrome(3))  // Output: [1 2]
	fmt.Println(EvenOddPalindrome(12)) // Output: [4 6]
}
