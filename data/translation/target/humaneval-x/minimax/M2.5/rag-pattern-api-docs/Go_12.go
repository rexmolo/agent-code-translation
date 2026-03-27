package main

import (
	"fmt"
	"slices"
)

func Longest(strings []string) interface{} {
	if len(strings) == 0 {
		return nil
	}

	// Find the maximum length
	lengths := make([]int, len(strings))
	for i, s := range strings {
		lengths[i] = len(s)
	}
	maxLen := slices.Max(lengths)

	// Return the first string with max length
	for _, s := range strings {
		if len(s) == maxLen {
			return s
		}
	}

	return nil
}

func main() {
	// Test cases
	fmt.Println(Longest([]string{}))
	fmt.Println(Longest([]string{"a", "b", "c"}))
	fmt.Println(Longest([]string{"a", "bb", "ccc"}))
}