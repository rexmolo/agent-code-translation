package main

import (
	"fmt"
	"strings"
)

func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")
	lastWord := parts[len(parts)-1]

	if len(lastWord) == 1 {
		lower := strings.ToLower(lastWord)
		c := lower[0]
		return c >= 'a' && c <= 'z'
	}
	return false
}

func main() {
	fmt.Println(CheckIfLastCharIsALetter("apple pie"))
	fmt.Println(CheckIfLastCharIsALetter("apple pi e"))
	fmt.Println(CheckIfLastCharIsALetter("apple pi e "))
	fmt.Println(CheckIfLastCharIsALetter(""))
}
