package main

import "strings"

func Concatenate(strings []string) string {
	return strings.Join(strings, "")
}

func main() {
	// Test cases to verify the function works
	// Empty list
	println(Concatenate([]string{}))
	// List with elements
	println(Concatenate([]string{"a", "b", "c"}))
}
