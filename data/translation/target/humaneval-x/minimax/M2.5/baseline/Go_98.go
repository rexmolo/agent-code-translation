package main

import (
	"fmt"
	"strings"
)

func CountUpper(s string) int {
	count := 0
	for i := 0; i < len(s); i += 2 {
		if strings.Contains("AEIOU", string(s[i])) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountUpper("aBCdEf")) // Output: 1
	fmt.Println(CountUpper("abcdefg")) // Output: 0
	fmt.Println(CountUpper("dBBE"))   // Output: 0
}
