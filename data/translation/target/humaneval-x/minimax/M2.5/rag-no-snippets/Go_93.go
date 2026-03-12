package main

import (
	"fmt"
	"strings"
)

func Encode(message string) string {
	vowels := "aeiouAEIOU"
	vowelsReplace := make(map[rune]rune)
	for _, v := range vowels {
		vowelsReplace[v] = v + 2
	}
	
	// Swap case: lowercase -> uppercase, uppercase -> lowercase
	message = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' {
			return r - 32 // convert to uppercase
		} else if r >= 'A' && r <= 'Z' {
			return r + 32 // convert to lowercase
		}
		return r
	}, message)
	
	// Replace vowels with the letter 2 positions ahead
	result := []rune{}
	for _, c := range message {
		if newChar, ok := vowelsReplace[c]; ok {
			result = append(result, newChar)
		} else {
			result = append(result, c)
		}
	}
	
	return string(result)
}

func main() {
	fmt.Println(Encode("test"))
	fmt.Println(Encode("This is a message"))
}
