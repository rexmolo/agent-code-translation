package main

import (
	"fmt"
	"strconv"
)

func CircularShift(x int, shift int) string {
	s := strconv.Itoa(x)
	if shift > len(s) {
		// Reverse the string
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	return s[len(s)-shift:] + s[:len(s)-shift]
}

func main() {
	// Test cases
	fmt.Println(CircularShift(12, 1)) // "21"
	fmt.Println(CircularShift(12, 2)) // "12"
}