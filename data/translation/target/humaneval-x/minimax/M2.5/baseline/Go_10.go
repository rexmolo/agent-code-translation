package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
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

	prefix := str[:beginningOfSuffix]
	reversedPrefix := reverseString(prefix)

	return str + reversedPrefix
}

func main() {
	// Test cases from the docstring
	fmt.Println(MakePalindrome(""))    // ""
	fmt.Println(MakePalindrome("cat"))  // "catac"
	fmt.Println(MakePalindrome("cata")) // "catac"
}