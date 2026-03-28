package main

import "fmt"

func isPalindrome(str string) bool {
	// Test if given string is a palindrome
	return str == reverseString(str)
}

func reverseString(str string) string {
	// Reverse a string
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func MakePalindrome(str string) string {
	// Find the shortest palindrome that begins with a supplied string.
	// Algorithm idea is simple:
	// - Find the longest postfix of supplied string that is a palindrome.
	// - Append to the end of the string reverse of a string prefix that comes before the palindromic suffix.
	if len(str) == 0 {
		return ""
	}

	beginningOfSuffix := 0

	for !isPalindrome(str[beginningOfSuffix:]) {
		beginningOfSuffix++
	}

	return str + reverseString(str[:beginningOfSuffix])
}

func main() {
	// Test cases from docstrings
	fmt.Println(MakePalindrome(""))    // ''
	fmt.Println(MakePalindrome("cat")) // catac
	fmt.Println(MakePalindrome("cata")) // catac
}
