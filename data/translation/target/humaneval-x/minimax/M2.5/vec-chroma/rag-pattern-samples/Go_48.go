package main

import "fmt"

func IsPalindrome(text string) bool {
	for i := 0; i < len(text); i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(IsPalindrome(""))      // true
	fmt.Println(IsPalindrome("aba"))   // true
	fmt.Println(IsPalindrome("aaaaa")) // true
	fmt.Println(IsPalindrome("zbcd"))  // false
}