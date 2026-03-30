package main

import "fmt"

func EvenOddPalindrome(n int) [2]int {
	isPalindrome := func(num int) bool {
		s := fmt.Sprintf("%d", num)
		// Check palindrome by comparing characters from both ends
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

tevenCount := 0
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