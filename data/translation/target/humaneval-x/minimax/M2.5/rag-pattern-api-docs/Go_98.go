package main

import (
	"strings"
)

func CountUpper(s string) int {
	count := 0
	for i := 0; i < len(s); i += 2 {
		if strings.ContainsRune("AEIOU", rune(s[i])) {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	println(CountUpper("aBCdEf")) // Expected: 1
	println(CountUpper("abcdefg")) // Expected: 0
	println(CountUpper("dBBE")) // Expected: 0
}