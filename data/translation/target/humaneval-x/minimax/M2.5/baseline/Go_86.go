package main

import (
	"fmt"
	"sort"
	"strings"
)

func AntiShuffle(s string) string {
	parts := strings.Split(s, " ")
	for i, part := range parts {
		runes := []rune(part)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		parts[i] = string(runes)
	}
	return strings.Join(parts, " ")
}

func main() {
	fmt.Println(AntiShuffle("Hi"))
	fmt.Println(AntiShuffle("hello"))
	fmt.Println(AntiShuffle("Hello World!!!"))
}
