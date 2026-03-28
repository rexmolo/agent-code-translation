package main

import (
	"fmt"
	"strings"
)

func SelectWords(s string, n int) []string {
	vowels := "aeiouAEIOU"
	result := []string{}
	words := strings.Fields(s)

	for _, word := range words {
		nConsonants := 0
		for _, ch := range word {
			if !strings.ContainsRune(vowels, ch) {
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
}
