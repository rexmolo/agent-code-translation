package main

import (
	"strings"
	"unicode"
)

func Digitsum(x string) int {
	if x == "" {
		return 0
	}

	sum := 0
	for _, char := range x {
		if unicode.IsUpper(char) {
			sum += int(char)
		}
	}
	return sum
}

func main() {
	// Test cases
	println(Digitsum(""))        // 0
	println(Digitsum("abAB"))    // 131
	println(Digitsum("abcCd"))   // 67
	println(Digitsum("helloE"))  // 69
	println(Digitsum("woArBld")) // 131
	println(Digitsum("aAaaaXa")) // 153
}