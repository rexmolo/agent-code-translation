package main

import (
	"fmt"
	"slices"
)

func Unique(l []int) []int {
	// Use map to get unique elements
	seen := make(map[int]struct{})
	for _, v := range l {
		seen[v] = struct{}{}
	}

	// Collect unique elements into a slice
	result := make([]int, 0, len(seen))
	for k := range seen {
		result = append(result, k)
	}

	// Sort the slice
	slices.Sort(result)

	return result
}

func main() {
	// Test the function
	result := Unique([]int{5, 3, 5, 2, 3, 3, 9, 0, 123})
	fmt.Println(result)
}