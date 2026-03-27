package main

import (
	"fmt"
)

func isPalindrome(str string) bool {
	runes := []rune(str)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func MakePalindrome(str string) string {
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
	fmt.Println(MakePalindrome(""))
	fmt.Println(MakePalindrome("cat"))
	fmt.Println(MakePalindrome("cata"))
}