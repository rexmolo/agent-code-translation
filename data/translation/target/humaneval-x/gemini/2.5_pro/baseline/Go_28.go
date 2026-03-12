package main

import (
	"fmt"
	"strings"
)

// Concatenate joins a slice of strings into a single string.
func Concatenate(stringsToJoin []string) string {
	return strings.Join(stringsToJoin, "")
}

func main() {
	// Example 1: equivalent to concatenate([])
	result1 := Concatenate([]string{})
	fmt.Printf("Result for []: %q\n", result1)

	// Example 2: equivalent to concatenate(['a', 'b', 'c'])
	result2 := Concatenate([]string{"a", "b", "c"})
	fmt.Printf("Result for [\"a\", \"b\", \"c\"]: %q\n", result2)
}
