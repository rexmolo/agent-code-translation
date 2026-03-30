package main

import (
	"fmt"
	"strings"
)

// IsPalindrome tests if given string is a palindrome
func IsPalindrome(str string) bool {
	return str == reverseString(str)
}

// reverseString returns the reverse of a string
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// MakePalindrome finds the shortest palindrome that begins with a supplied string.
// Algorithm idea is simple:
// - Find the longest postfix of supplied string that is a palindrome.
// - Append to the end of the string reverse of a string prefix that comes before the palindromic suffix.
func MakePalindrome(str string) string {
	if len(str) == 0 {
		return ""
	}

	beginningOfSuffix := 0

	for !IsPalindrome(str[beginningOfSuffix:]) {
		beginningOfSuffix++
	}

	return str + reverseString(str[:beginningOfSuffix])
}

func main() {
	// Test cases from docstring
	fmt.Println(MakePalindrome(""))      // ""
	fmt.Println(MakePalindrome("cat"))  // "catac"
	fmt.Println(MakePalindrome("cata")) // "catac"
}