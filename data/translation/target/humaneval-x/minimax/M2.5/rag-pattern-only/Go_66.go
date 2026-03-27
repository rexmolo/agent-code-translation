package main

import (
	"fmt"
	"unicode"
)

func Digitsum(s string) int {
	if s == "" {
		return 0
	}

	sum := 0
	for _, r := range s {
		if unicode.IsUpper(r) {
			sum += int(r)
		}
	}

	return sum
}

func main() {
	fmt.Println(Digitsum(""))       // 0
	fmt.Println(Digitsum("abAB"))   // 131
	fmt.Println(Digitsum("abcCd"))  // 67
	fmt.Println(Digitsum("helloE")) // 69
	fmt.Println(Digitsum("woArBld")) // 131
	fmt.Println(Digitsum("aAaaaXa")) // 153
}
