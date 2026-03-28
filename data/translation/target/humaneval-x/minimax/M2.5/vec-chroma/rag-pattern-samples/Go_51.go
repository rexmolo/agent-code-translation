package main

import (
	"strings"
)

func RemoveVowels(text string) string {
	var result []rune
	for _, r := range text {
		if !strings.Contains("aeiou", strings.ToLower(string(r))) {
			result = append(result, r)
		}
	}
	return string(result)
}
