package main

import "fmt"

// isPalindrome checks if a string is a palindrome
func isPalindrome(str string) bool {
	// Convert to runes to handle Unicode properly
	runes := []rune(str)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

// reverseString reverses a string
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// MakePalindrome finds the shortest palindrome that begins with a supplied string.
// Algorithm: Find the longest postfix of supplied string that is a palindrome.
// Append to the end of the string reverse of a string prefix that comes before the palindromic suffix.
func MakePalindrome(str string) string {
	if str == "" {
		return ""
	}

	beginningOfSuffix := 0

	for !isPalindrome(str[beginningOfSuffix:]) {
		beginningOfSuffix++
	}

	prefix := str[:beginningOfSuffix]
	reversedPrefix := reverseString(prefix)

	return str + reversedPrefix
}

func main() {
	fmt.Println(MakePalindrome(""))
	fmt.Println(MakePalindrome("cat"))
	fmt.Println(MakePalindrome("cata"))
}
