package main

import (
	"fmt"
	"strings"
	"unicode"
)

func SelectWords(s string, n int) []string {
	result := []string{}
	words := strings.Fields(s)

	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	for _, word := range words {
		nConsonants := 0
		for _, char := range word {
			lowerChar := unicode.ToLower(char)
			if !vowels[lowerChar] {
				nConsonants++
			}
		}
		if nConsonants == n {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	fmt.Println(SelectWords("Mary had a little lamb", 4))
	fmt.Println(SelectWords("Mary had a little lamb", 3))
	fmt.Println(SelectWords("simple white space", 2))
	fmt.Println(SelectWords("Hello world", 4))
	fmt.Println(SelectWords("Uncle sam", 3))
	fmt.Println(SelectWords("", 3))
}