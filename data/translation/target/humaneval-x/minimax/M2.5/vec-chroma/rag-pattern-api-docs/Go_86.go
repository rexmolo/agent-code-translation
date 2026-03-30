package main

import (
	"fmt"
	"sort"
	"strings"
)

func AntiShuffle(s string) string {
	words := strings.Split(s, " ")
	var result []string

	for _, word := range words {
		runes := []rune(word)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		result = append(result, string(runes))
	}

	return strings.Join(result, " ")
}

func main() {
	fmt.Println(AntiShuffle("Hi"))
	fmt.Println(AntiShuffle("hello"))
	fmt.Println(AntiShuffle("Hello World!!!"))
}
