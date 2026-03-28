package main

import "strings"

func Concatenate(strings []string) string {
	return strings.Join(strings, "")
}

func main() {
	// Example usage
	result1 := Concatenate([]string{})
	println(result1) // Output: (empty)
	
	result2 := Concatenate([]string{"a", "b", "c"})
	println(result2) // Output: abc
}