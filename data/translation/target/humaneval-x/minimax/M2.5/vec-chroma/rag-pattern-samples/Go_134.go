package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")
	check := parts[len(parts)-1]

	if len(check) == 1 && unicode.IsLetter(rune(check[0])) {
		return true
	}
	return false
}

func main() {
	fmt.Println(CheckIfLastCharIsALetter("apple pie"))   // false
	fmt.Println(CheckIfLastCharIsALetter("apple pi e"))  // true
	fmt.Println(CheckIfLastCharIsALetter("apple pi e ")) // false
	fmt.Println(CheckIfLastCharIsALetter(""))            // false
}