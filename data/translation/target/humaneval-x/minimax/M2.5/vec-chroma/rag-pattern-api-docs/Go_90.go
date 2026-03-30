package main

import (
	"fmt"
	"sort"
)

func NextSmallest(lst []int) interface{} {
	if len(lst) < 2 {
		return nil
	}

	// Create a set using map[int]struct{} to get unique elements
	seen := make(map[int]struct{})
	unique := []int{}

	for _, v := range lst {
		if _, exists := seen[v]; !exists {
			seen[v] = struct{}{}
			unique = append(unique, v)
		}
	}

	// Sort the unique elements in ascending order
	sort.Ints(unique)

	// Return second smallest if it exists
	if len(unique) < 2 {
		return nil
	}
	return unique[1]
}

func main() {
	// Test cases
	tests := [][]int{
		{1, 2, 3, 4, 5},
		{5, 1, 4, 3, 2},
		{},
		{1, 1},
	}

	for _, test := range tests {
		result := NextSmallest(test)
		fmt.Printf("NextSmallest(%v) = %v\n", test, result)
	}
}
