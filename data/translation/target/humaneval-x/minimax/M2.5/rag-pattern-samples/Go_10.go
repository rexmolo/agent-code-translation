package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	return str == reverseString(str)
}

func reverseString(str string) string {
	runes := []rune(str)
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
	// Test cases
	fmt.Println(MakePalindrome(""))    // Expected: ""
	fmt.Println(MakePalindrome("cat")) // Expected: "catac"
	fmt.Println(MakePalindrome("cata")) // Expected: "catac"
}
