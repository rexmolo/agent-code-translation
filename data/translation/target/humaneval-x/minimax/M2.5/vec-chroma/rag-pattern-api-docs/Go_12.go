package main

import (
	"fmt"
	"slices"
)

func Longest(strings []string) interface{} {
	if len(strings) == 0 {
		return nil
	}

	// Get the maximum length using slices.Max with a transformed slice
	maxLen := slices.MaxFunc(strings, func(a, b string) int {
		return len(a) - len(b)
	})
	_ = maxLen // we only need its length, not the value

	// Find the actual max length
	maxLen = 0
	for _, s := range strings {
		if len(s) > maxLen {
			maxLen = len(s)
		}
	}

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