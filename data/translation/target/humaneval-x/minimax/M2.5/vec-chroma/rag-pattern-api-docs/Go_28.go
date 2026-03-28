package main

import "strings"

// Concatenate concatenates a list of strings into a single string.
func Concatenate(strings []string) string {
	return strings.Join(strings, "")
}

func main() {
	// Test cases
	empty := Concatenate([]string{})
	result := Concatenate([]string{"a", "b", "c"})

	// Print results to verify
	println("Empty:", empty)
	println("Result:", result)
}