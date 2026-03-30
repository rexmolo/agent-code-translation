package main

import (
	"fmt"
	"unicode"
)

func Digitsum(x string) int {
	if x == "" {
		return 0
	}
	sum := 0
	for _, r := range x {
		if unicode.IsUpper(r) {
			sum += int(r)
		}
	}
	return sum
}

func main() {
	// Test examples
	fmt.Println(Digitsum(""))       // 0
	fmt.Println(Digitsum("abAB"))   // 131
	fmt.Println(Digitsum("abcCd"))  // 67
	fmt.Println(Digitsum("helloE")) // 69
	fmt.Println(Digitsum("woArBld"))// 131
	fmt.Println(Digitsum("aAaaaXa"))// 153
}
