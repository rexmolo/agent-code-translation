package main

import "unicode"

func Digitsum(x string) int {
	totalSum := 0
	for _, char := range x {
		if unicode.IsUpper(char) {
			totalSum += int(char)
		}
	}
	return totalSum
}
