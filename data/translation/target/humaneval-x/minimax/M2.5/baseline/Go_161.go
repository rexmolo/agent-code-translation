package main

import (
	"unicode"
)

func Solve(s string) string {
	runes := []rune(s)
	flg := false

	for i, r := range runes {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				runes[i] = unicode.ToLower(r)
			} else {
				runes[i] = unicode.ToUpper(r)
			}
			flg = true
		}
	}

	if !flg {
		// Reverse the string
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
	}

	return string(runes)
}