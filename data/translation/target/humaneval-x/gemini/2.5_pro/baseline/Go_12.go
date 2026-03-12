package main

import (
	"fmt"
)

// Longest finds the longest string in a slice of strings.
// It returns the first longest string found in case of a tie.
// It returns nil if the input slice is empty, matching Python's None.
func Longest(strings []string) interface{} {
	if len(strings) == 0 {
		return nil
	}

	maxLength := 0
	for _, s := range strings {
		if len(s) > maxLength {
			maxLength = len(s)
		}
	}

	for _, s := range strings {
		if len(s) == maxLength {
			return s
		}
	}

	return nil // Logically unreachable
}

func main() {
	// Replicating the Python docstring examples

	// >>> longest([])
	fmt.Printf("%#v\n", Longest([]string{}))

	// >>> longest(['a', 'b', 'c'])
	// 'a'
	fmt.Printf("%#v\n", Longest([]string{"a", "b", "c"}))

	// >>> longest(['a', 'bb', 'ccc'])
	// 'ccc'
	fmt.Printf("%#v\n", Longest([]string{"a", "bb", "ccc"}))
}
