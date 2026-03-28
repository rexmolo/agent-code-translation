package main

import (
	"fmt"
	"strings"
)

func CycpatternCheck(a, b string) bool {
	l := len(b)
	if l == 0 {
		return true
	}
	if l > len(a) {
		return false
	}

	pat := b + b

	// Check if b is directly a substring of a
	if strings.Contains(a, b) {
		return true
	}

	// Check all rotations of b
	for j := 0; j < l; j++ {
		rotation := pat[j : j+l]
		if strings.Contains(a, rotation) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(CycpatternCheck("abcd", "abd"))   // False
	fmt.Println(CycpatternCheck("hello", "ell")) // True
	fmt.Println(CycpatternCheck("whassup", "psus")) // False
	fmt.Println(CycpatternCheck("abab", "baa")) // True
	fmt.Println(CycpatternCheck("efef", "eeff")) // False
	fmt.Println(CycpatternCheck("himenss", "simen")) // True
}
