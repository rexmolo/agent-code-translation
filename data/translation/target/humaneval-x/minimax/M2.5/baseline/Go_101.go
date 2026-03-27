package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	// Replace commas with spaces
	s = strings.ReplaceAll(s, ",", " ")

	// Split by whitespace (handles multiple spaces)
	return strings.Fields(s)
}

func main() {
	// Test cases
	test1 := "Hi, my name is John"
	test2 := "One, two, three, four, five, six"

	result1 := WordsString(test1)
	result2 := WordsString(test2)

	// Print results
	for _, w := range result1 {
		println(w)
	}
	for _, w := range result2 {
		println(w)
	}
}