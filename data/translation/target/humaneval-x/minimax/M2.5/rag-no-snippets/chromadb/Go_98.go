package main

import (
	"fmt"
	"strings"
)

func CountUpper(s string) int {
	vowels := "AEIOU"
	count := 0
	for i := 0; i < len(s); i += 2 {
		if strings.Contains(vowels, string(s[i])) {
			count++
		}
	}
	return count
}

func main() {
	// Test cases
	fmt.Println(CountUpper("aBCdEf")) // 1
	fmt.Println(CountUpper("abcdefg")) // 0
	fmt.Println(CountUpper("dBBE")) // 0
}