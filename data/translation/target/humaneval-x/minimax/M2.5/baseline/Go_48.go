package main

import "fmt"

func IsPalindrome(text string) bool {
	n := len(text)
	for i := 0; i < n/2; i++ {
		if text[i] != text[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases from docstrings
	fmt.Println(IsPalindrome(""))      // true
	fmt.Println(IsPalindrome("aba"))   // true
	fmt.Println(IsPalindrome("aaaaa")) // true
	fmt.Println(IsPalindrome("zbcd"))  // false
}