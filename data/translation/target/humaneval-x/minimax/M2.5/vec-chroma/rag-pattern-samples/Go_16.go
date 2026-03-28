package main

import (
	"fmt"
	"strings"
)

func CountDistinctCharacters(str string) int {
	str = strings.ToLower(str)
	unique := make(map[rune]bool)
	for _, c := range str {
		unique[c] = true
	}
	return len(unique)
}

func main() {
	fmt.Println(CountDistinctCharacters("xyzXYZ"))
	fmt.Println(CountDistinctCharacters("Jerry"))
}