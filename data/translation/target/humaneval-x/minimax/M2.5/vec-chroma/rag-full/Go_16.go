package main

import (
	"strings"
)

func CountDistinctCharacters(str string) int {
	lower := strings.ToLower(str)
	unique := make(map[rune]struct{})
	for _, r := range lower {
		unique[r] = struct{}{}
	}
	return len(unique)
}