package main

import "fmt"

func IsPalindrome(text string) bool {
	runes := []rune(text)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		if runes[i] != runes[length-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases
	fmt.Println(IsPalindrome(""))    // true
	fmt.Println(IsPalindrome("aba")) // true
	fmt.Println(IsPalindrome("aaaaa")) // true
	fmt.Println(IsPalindrome("zbcd")) // false
}
