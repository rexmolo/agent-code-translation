package main

import (
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