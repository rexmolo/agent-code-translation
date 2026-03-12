package main

import (
	"fmt"
	"unicode"
)

func Solve(s string) string {
	flg := 0
	newStr := []rune(s)

	for idx, r := range newStr {
		if unicode.IsLetter(r) {
			if unicode.IsLower(r) {
				newStr[idx] = unicode.ToUpper(r)
			} else {
				newStr[idx] = unicode.ToLower(r)
			}
			flg = 1
		}
	}

	result := string(newStr)

	if flg == 0 {
		// Reverse the string
		runes := []rune(result)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	return result
}

func main() {
	// Test cases
	testCases := []string{"1234", "ab", "#a@C"}
	for _, tc := range testCases {
		fmt.Println(Solve(tc))
	}
}