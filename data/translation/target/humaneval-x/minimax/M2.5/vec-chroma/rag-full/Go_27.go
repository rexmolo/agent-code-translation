package main

import (
	"fmt"
	"strings"
)

func FlipCase(str string) string {
	return strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r + 32 // Convert uppercase to lowercase
		}
		if r >= 'a' && r <= 'z' {
			return r - 32 // Convert lowercase to uppercase
		}
		return r // Keep other characters unchanged
	}, str)
}

func main() {
	// Test the FlipCase function
	testCases := []string{"Hello", "World", "PyThOn", "123ABC!"}
	for _, tc := range testCases {
		fmt.Printf("%s -> %s\n", tc, FlipCase(tc))
	}
}