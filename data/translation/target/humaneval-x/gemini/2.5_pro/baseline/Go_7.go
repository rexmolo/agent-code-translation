package main

import (
	"fmt"
	"strings"
)

// FilterBySubstring filters an input slice of strings, keeping only the ones
// that contain the given substring.
func FilterBySubstring(stringList []string, substring string) []string {
	var result []string
	for _, s := range stringList {
		if strings.Contains(s, substring) {
			result = append(result, s)
		}
	}
	return result
}

// main function to demonstrate the usage of FilterBySubstring
// It mimics the doctests in the original Python code.
func main() {
	// Test case 1: filter_by_substring([], 'a') -> []
	list1 := []string{}
	result1 := FilterBySubstring(list1, "a")
	fmt.Printf("Input: %v, Substring: 'a', Result: %v\n", list1, result1)

	// Test case 2: filter_by_substring(['abc', 'bacd', 'cde', 'array'], 'a') -> ['abc', 'bacd', 'array']
	list2 := []string{"abc", "bacd", "cde", "array"}
	result2 := FilterBySubstring(list2, "a")
	fmt.Printf("Input: %v, Substring: 'a', Result: %v\n", list2, result2)
}
