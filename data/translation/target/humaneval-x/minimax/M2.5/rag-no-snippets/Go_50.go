package main

import (
	"fmt"
	"strings"
)

func DecodeShift(s string) string {
	result := make([]rune, len(s))
	for i, ch := range s {
		// Shift character back by 5, handling wraparound
		result[i] = rune(((int(ch) - 5 - int('a') + 26) % 26) + int('a'))
	}
	return string(result)
}

func main() {
	// Test example
	encoded := "fghjkl" // encoded version of "abcde"
	decoded := DecodeShift(encoded)
	fmt.Println(decoded)
}
