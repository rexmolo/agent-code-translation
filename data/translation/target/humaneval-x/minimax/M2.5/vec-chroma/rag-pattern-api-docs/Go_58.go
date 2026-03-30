package main

import (
	"slices"
)

func Common(l1 []int, l2 []int) []int {
	// Create a map for O(1) lookup of elements in l1
	inL1 := make(map[int]bool)
	for _, e := range l1 {
		inL1[e] = true
	}

	// Find common elements from l2, tracking which we've already added
	found := make(map[int]bool)
	var result []int
	for _, e := range l2 {
		if inL1[e] && !found[e] {
			result = append(result, e)
			found[e] = true
		}
	}

	// Sort the result in ascending order
	slices.Sort(result)

	return result
}

func main() {
	// Example usage
	result := Common([]int{1, 4, 3, 34, 653, 2, 5}, []int{5, 7, 1, 5, 9, 653, 121})
	println(result) // Should print: [1 5 653]
}
