package main

import (
	"fmt"
	"sort"
)

func Unique(l []int) []int {
	// Use a map to track seen elements (simulating a set)
	seen := make(map[int]bool)

	// Collect unique elements into a slice
	unique := make([]int, 0)
	for _, v := range l {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

	// Sort the unique elements
	sort.Ints(unique)

	return unique
}

func main() {
	// Test the function with the example from docstring
	result := Unique([]int{5, 3, 5, 2, 3, 3, 9, 0, 123})
	fmt.Println(result)
}
