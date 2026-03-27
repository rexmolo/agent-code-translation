package main

import "strings"

// Concatenate concatenates a list of strings into a single string
func Concatenate(strings []string) string {
	return strings.Join(strings, "")
}

func main() {
	// Example usage - can be tested by running the program
	_ = Concatenate([]string{})
	_ = Concatenate([]string{"a", "b", "c"})
}