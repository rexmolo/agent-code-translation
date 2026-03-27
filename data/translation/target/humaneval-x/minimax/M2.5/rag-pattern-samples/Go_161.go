package main

import (
	"unicode"
)

func Solve(s string) string {
	flg := 0
	runes := []rune(s)

	for i, c := range runes {
		if unicode.IsLetter(c) {
			if unicode.IsUpper(c) {
				runes[i] = unicode.ToLower(c)
			} else {
				runes[i] = unicode.ToUpper(c)
			}
			flg = 1
		}
	}

	if flg == 0 {
		// Reverse the string
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}

	return string(runes)
}
