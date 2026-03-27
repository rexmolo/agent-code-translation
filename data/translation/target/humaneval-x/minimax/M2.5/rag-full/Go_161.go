package main

import (
	"fmt"
	"unicode"
)

func Solve(s string) string {
	flg := 0
	runes := []rune(s)
	newStr := make([]rune, len(runes))

	idx := 0
	for _, r := range runes {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				newStr[idx] = unicode.ToLower(r)
			} else {
				newStr[idx] = unicode.ToUpper(r)
			}
			flg = 1
		} else {
			newStr[idx] = r
		}
		idx++
	}

	if flg == 0 {
		// Reverse the string
		for i, j := 0, len(newStr)-1; i < j; i, j = i+1, j-1 {
			newStr[i], newStr[j] = newStr[j], newStr[i]
		}
	}

	return string(newStr)
}

func main() {
	// Test cases
	fmt.Println(Solve("1234")) // "4321"
	fmt.Println(Solve("ab"))  // "AB"
	fmt.Println(Solve("#a@C")) // "#A@c"
}