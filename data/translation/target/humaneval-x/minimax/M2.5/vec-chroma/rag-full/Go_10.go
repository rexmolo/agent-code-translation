package main

import (
	"fmt"
	"strings"
)

// isPalindrome tests if a given string is a palindrome
func isPalindrome(str string) bool {
	return str == reverseString(str)
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// MakePalindrome finds the shortest palindrome that begins with a supplied string.
// Algorithm:
// - Find the longest postfix of supplied string that is a palindrome.
// - Append to the end of the string reverse of a string prefix that comes before the palindromic suffix.
func MakePalindrome(str string) string {
	if str == "" {
		return ""
	}

	beginningOfSuffix := 0

	for !isPalindrome(str[beginningOfSuffix:]) {
		beginningOfSuffix++
	}

	return str + reverseString(str[:beginningOfSuffix])
}

func main() {
	// Test cases from docstring
	fmt.Println(MakePalindrome(""))
	fmt.Println(MakePalindrome("cat"))
	fmt.Println(MakePalindrome("cata"))
}