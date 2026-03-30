package main

import (
	"fmt"
	"strings"
	"unicode"
)

func SelectWords(s string, n int) []string {
	vowels := map[rune]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	}

	result := []string{}

	for _, word := range strings.Fields(s) {
		consonantCount := 0
		for _, ch := range word {
			lower := unicode.ToLower(ch)
			if !vowels[lower] {
				consonantCount++
			}
		}
		if consonantCount == n {
			result = append(result, word)
		}
	}

	return result
}

func main() {
	fmt.Println(SelectWords("Mary had a little lamb", 4)) // [little]
	fmt.Println(SelectWords("Mary had a little lamb", 3)) // [Mary lamb]
	fmt.Println(SelectWords("simple white space", 2))     // []
	fmt.Println(SelectWords("Hello world", 4))            // [world]
	fmt.Println(SelectWords("Uncle sam", 3))              // [Uncle]
}
