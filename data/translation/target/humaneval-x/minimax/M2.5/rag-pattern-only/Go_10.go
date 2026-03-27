package main

import "fmt"

func isPalindrome(s string) bool {
	// Reverse the string and compare
	reversed := reverseString(s)
	return s == reversed
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

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
	fmt.Println(MakePalindrome(""))    // ''
	fmt.Println(MakePalindrome("cat")) // catac
	fmt.Println(MakePalindrome("cata")) // catac
}