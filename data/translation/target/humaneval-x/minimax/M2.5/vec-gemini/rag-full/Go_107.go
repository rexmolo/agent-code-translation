package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EvenOddPalindrome(n int) [2]int {
	isPalindrome := func(num int) bool {
		s := strconv.Itoa(num)
		var b strings.Builder
		for i := len(s) - 1; i >= 0; i-- {
			b.WriteByte(s[i])
		}
		return s == b.String()
	}

devenPalindromeCount := 0
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
