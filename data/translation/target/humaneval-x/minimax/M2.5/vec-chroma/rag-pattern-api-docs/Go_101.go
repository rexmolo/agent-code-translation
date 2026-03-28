package main

import (
	"strings"
)

func WordsString(s string) []string {
	if s == "" {
		return []string{}
	}

	s = strings.ReplaceAll(s, ",", " ")
	return strings.Fields(s)
}

func main() {
	// Test cases
	input1 := "Hi, my name is John"
	result1 := WordsString(input1)
	// Expected: ["Hi", "my", "name", "is", "John"]
	println(result1)

	input2 := "One, two, three, four, five, six"
	result2 := WordsString(input2)
	// Expected: ["One", "two", "three", "four", "five", "six"]
	println(result2)
}