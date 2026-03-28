package main

import (
	"strings"
)

// Concatenate concatenates a slice of strings into a single string.
func Concatenate(strings []string) string {
	return strings.Join(strings, "")
}

// For testing purposes
func main() {
	// Test cases
	var empty []string
	printResult(empty)
	printResult([]string{"a", "b", "c"})
}

func printResult(strings []string) {
	result := Concatenate(strings)
	println(result)
}