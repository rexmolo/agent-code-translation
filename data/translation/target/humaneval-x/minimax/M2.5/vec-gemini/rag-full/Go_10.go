package main

import "fmt"

// isPalindrome checks if a string is a palindrome
func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

// MakePalindrome finds the shortest palindrome that begins with a supplied string.
// Algorithm idea is simple:
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

	// Reverse the prefix before the palindromic suffix
	prefix := str[:beginningOfSuffix]
	reversedPrefix := ""
	for i := len(prefix) - 1; i >= 0; i-- {
		reversedPrefix += string(prefix[i])
	}

	return str + reversedPrefix
}

func main() {
	// Test cases from docstrings
	fmt.Println(MakePalindrome(""))    // Expected: ""
	fmt.Println(MakePalindrome("cat")) // Expected: "catac"
	fmt.Println(MakePalindrome("cata")) // Expected: "catac"
}