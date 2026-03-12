package main

import "fmt"

func IsPalindrome(text string) bool {
	// Convert to runes to handle Unicode characters properly
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases from the docstring
	fmt.Println(IsPalindrome(""))      // true
	fmt.Println(IsPalindrome("aba"))   // true
	fmt.Println(IsPalindrome("aaaaa")) // true
	fmt.Println(IsPalindrome("zbcd"))  // false
}
