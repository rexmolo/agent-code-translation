package main

import (
	"strings"
	"unicode"
)

func CheckIfLastCharIsALetter(txt string) bool {
	parts := strings.Split(txt, " ")
	lastElement := parts[len(parts)-1]
	
	if len(lastElement) != 1 {
		return false
	}
	
	lowerChar := unicode.ToLower(rune(lastElement[0]))
	return lowerChar >= 'a' && lowerChar <= 'z'
}
