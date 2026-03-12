package main

import (
	"strings"
)

func CycpatternCheck(a, b string) bool {
	l := len(b)
	if l > len(a) {
		return false
	}
	pat := b + b
	for i := 0; i <= len(a)-l; i++ {
		for j := 0; j <= l; j++ {
			if a[i:i+l] == pat[j:j+l] {
				return true
			}
		}
	}
	return false
}

func main() {
	// Test cases
	println(CycpatternCheck("abcd", "abd"))    // false
	println(CycpatternCheck("hello", "ell"))   // true
	println(CycpatternCheck("whassup", "psus")) // false
	println(CycpatternCheck("abab", "baa"))   // true
	println(CycpatternCheck("efef", "eeff"))  // false
	println(CycpatternCheck("himenss", "simen")) // true
}
