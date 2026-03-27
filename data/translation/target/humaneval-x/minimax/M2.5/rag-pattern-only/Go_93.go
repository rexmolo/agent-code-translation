package main

import (
	"fmt"
	"strings"
)

func Encode(message string) string {
	vowels := "aeiouAEIOU"
	vowelReplacements := make(map[rune]rune, len(vowels))
	for _, v := range vowels {
		vowelReplacements[v] = v + 2
	}

	var result strings.Builder
	for _, char := range message {
		swapped := swapCase(char)
		if repl, ok := vowelReplacements[swapped]; ok {
			result.WriteRune(repl)
		} else {
			result.WriteRune(swapped)
		}
	}
	return result.String()
}

func swapCase(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

func main() {
	fmt.Println(Encode("test"))
	fmt.Println(Encode("This is a message"))
}