package main

import (
	"fmt"
	"strconv"
)

func EvenOddPalindrome(n int) [2]int {
	isPalindrome := func(num int) bool {
		s := strconv.Itoa(num)
		reversed := ""
		for i := len(s) - 1; i >= 0; i-- {
			reversed += string(s[i])
		}
		return s == reversed
	}

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
	fmt.Println(EvenOddPalindrome(3))
	fmt.Println(EvenOddPalindrome(12))
}